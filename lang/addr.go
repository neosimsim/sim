package lang

import (
	"bufio"
	"io"
	"regexp"
	"unicode"

	file "github.com/neosimsim/sim/file"
)

// Address is an interface to provide functions to scan a file for an
// Address.
type Address interface {
	// Scan for the Address in File described by AddrresDesc after Address start.
	Scan(f *file.File, dot file.Range) (file.Range, error)
}

func ParseAddress(r io.Reader) (Address, error) {
	var addr Address
	addr, err := parseComposed(r)
	if err != nil {
		return nil, err
	}
	bufReader := bufio.NewReader(r)
	addrSeparator, _, err := bufReader.ReadRune()
	if err == io.EOF {
		return addr, nil
	} else if err != nil {
		return nil, err
	}
	switch addrSeparator {
	case ',':
		endOfRange, err := ParseAddress(r)
		if err != nil {
			return nil, err
		}
		addr = &rangeAddress{Start: addr, End: endOfRange}
	case ';':
		endOfRange, err := ParseAddress(r)
		if err != nil {
			return nil, err
		}
		addr = &rangeAddress{Start: addr, End: endOfRange, relative: true}
	default:
		err := bufReader.UnreadRune()
		if err != nil {
			return nil, err
		}
	}
	return addr, nil
}

// Models a rangeAddress of addresses. A range is basically two
// addresses seperated by ',' or ';'. Note that the
// ',' and ';' composite have lower precedense than
// '+' or '-'.
// rangeAddresss are modeled a b-tree. Address compsitions with ; and ,
// has lower precedence. Hence other addres are leafes and ; and , are parentes
type rangeAddress struct {
	Start    Address
	End      Address
	relative bool
}

// composed addresses are separated by + and -
type composedAddress struct {
	this      Address
	Next      Address
	direction direction
}

func (r *rangeAddress) Scan(f *file.File, dot file.Range) (file.Range, error) {
	start, err := r.Start.Scan(f, dot)
	if err != nil {
		return file.Range{}, err
	}
	offset := dot
	if r.relative {
		offset = start
	}
	end, err := r.End.Scan(f, offset)
	if err != nil {
		return file.Range{}, err
	}
	return file.Range{Start: start.Start, End: end.End}, nil
}

func parseRange(r io.Reader) (*rangeAddress, error) {
	return nil, nil
}

func parseComposed(r io.Reader) (startOfComposition Address, err error) {
	bufReader := bufio.NewReader(r)
	startOfComposition, err = parseAddress(r)
	composedAddr := &composedAddress{this: startOfComposition, Next: nil, direction: forwards}
	previousAddr := composedAddr
	startOfComposition = composedAddr
	if err != nil {
		return nil, err
	}
	for {
		directionSet := false
		var next Address
		addrSeparator, _, err := bufReader.ReadRune()
		if err == io.EOF {
			return startOfComposition, nil
		} else if err != nil {
			return nil, err
		}
		direction := forwards
		switch addrSeparator {
		case '+': // read address
			directionSet = true
			next, err = parseAddress(r)
		case '-': // set direction backward, read address
			directionSet = true
			direction = backwards
			next, err = parseAddress(r)
		default:
			bufReader.UnreadRune()
			next, err = parseAddress(r)
		}
		if err != nil {
			return nil, err
		}
		if _, ok := next.(*dotAddress); ok {
			if startOfComposition == composedAddr { // only one address read
				return composedAddr.this, nil
			}
			if directionSet {
				// in case we know an addres has to follow, because '+' or '-' has been read.
				composedAddr.Next = lineAddress(1)
			} else {
				// We don't want a composotion to and with a composedAddress with Next == nil.
				previousAddr.Next = composedAddr.this
			}
			return startOfComposition, nil
		}
		previousAddr = composedAddr
		new := &composedAddress{this: next, Next: nil, direction: direction}
		composedAddr.Next = new
		composedAddr = new
	}
	return startOfComposition, nil
}

func parseAddress(r io.Reader) (addr Address, err error) {
	bufReader := bufio.NewReader(r)
	reader := NewLangReader(bufReader)
	addrMark, _, err := bufReader.ReadRune()
	if err == io.EOF {
		return &dotAddress{}, nil
	} else if err != nil {
		return nil, err
	}
	switch addrMark {
	case '"': // read file scanner regexp until '"' TODO
	case '#': // read number
		var number int
		number, err = reader.ReadNumber()
		addr = offsetAddress(number)
	case '.': // dot scanner
		addr = &dotAddress{}
	case '/': // read regexp until '/'
		var regexpString string
		regexpString, err = reader.ReadLineTo('/')
		addr = &regexpAddress{regexp: regexp.MustCompilePOSIX(regexpString), direction: forwards}
	case '?': // read regexp until '?', set direction backwards
		var regexpString string
		regexpString, err = reader.ReadLineTo('?')
		addr = &regexpAddress{regexp: regexp.MustCompilePOSIX(regexpString), direction: backwards}
	case '0': // looks like a line offset but is a rune offset 0
		number, _ := reader.ReadNumber()
		if number == 0 {
			addr = offsetAddress(number)
		} else {
			addr = lineAddress(number)
		}
	default: // expect a number as line offset
		err = bufReader.UnreadRune()
		if !unicode.IsDigit(addrMark) {
			addr = &dotAddress{}
			return
		}
		var number int
		number, err = reader.ReadNumber()
		addr = lineAddress(number)
	}
	return
}

type dotAddress struct{}

func (a *dotAddress) Scan(f *file.File, dot file.Range) (file.Range, error) {
	return file.Range{}, nil
}

type offsetAddress int

func (a offsetAddress) Scan(f *file.File, dot file.Range) (file.Range, error) {
	return file.Range{}, nil
}

type lineAddress int

func (a lineAddress) Scan(f *file.File, dot file.Range) (file.Range, error) {
	return file.Range{}, nil
}

type regexpAddress struct {
	regexp    *regexp.Regexp
	direction direction
}

func (a *regexpAddress) Scan(f *file.File, dot file.Range) (file.Range, error) {
	return file.Range{}, nil
}

type direction bool

func (d direction) String() string {
	if d == forwards {
		return "forwards"
	}
	return "backwards"
}

const (
	forwards  direction = true
	backwards direction = false
)

func (a *composedAddress) Scan(f *file.File, dot file.Range) (file.Range, error) {
	start, err := a.this.Scan(f, dot)
	if err != nil {
		return file.Range{}, err
	}
	end, err := a.Next.Scan(f, dot)
	if err != nil {
		return file.Range{}, err
	}
	return file.Range{Start: start.Start, End: end.End}, nil
}
