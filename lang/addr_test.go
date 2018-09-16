package lang

import (
	"bufio"
	"io"
	"log"
	"strings"
	"testing"
)

func TestParseRangeAddress(t *testing.T) {
	bufReader := bufio.NewReader(strings.NewReader("1,2d"))
	a, err := ParseAddress(bufReader)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if addr, ok := a.(*rangeAddress); !ok {
		t.Fatalf("unexpected address type, got %T", a)
	} else {
		if addr.relative {
			t.Fatalf("expected range address to be absolute")
		}
	}
	rest, _, err := bufReader.ReadRune()
	if err != nil {
		t.Fatalf("unexpected error reading rest %v", err)
	}
	if rest != 'd' {
		t.Fatalf("unexpected rest %s", string(rest))
	}

	bufReader = bufio.NewReader(strings.NewReader("1;2"))
	a, err = ParseAddress(bufReader)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if addr, ok := a.(*rangeAddress); !ok {
		t.Fatalf("unexpected address type, got %T", a)
	} else {
		if !addr.relative {
			t.Fatalf("expected range address to be relative")
		}
	}
	rest, _, err = bufReader.ReadRune()
	if err != io.EOF {
		t.Fatalf("unexpected error reading rest %v", err)
	}
}

func TestParseRangeAddress2(t *testing.T) {
	bufReader := bufio.NewReader(strings.NewReader("3,4/hello/c/hallo"))
	a, err := ParseAddress(bufReader)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if addr, ok := a.(*rangeAddress); !ok {
		t.Fatalf("unexpected address type, got %T", a)
	} else {
		if addr.relative {
			t.Fatalf("expected range address to be absolute")
		}
		if start, ok := addr.Start.(lineAddress); !ok {
			t.Fatalf("unexpected address type, got %T", a)
		} else if start != lineAddress(3) {
			t.Fatalf("expected line address at start %v", start)
		}
		if end, ok := addr.End.(*composedAddress); !ok {
			t.Fatalf("unexpected address type, got %T %v", addr.End, addr.End)
		} else {
			if lineAddress2, ok := end.this.(lineAddress); !ok {
				t.Fatalf("unexpected address type, got %T %v", end.this, end.this)
			} else if lineAddress2 != lineAddress(4) {
				t.Fatalf("expected line address 2 %v", lineAddress2)
			}
			if _, ok := end.Next.(*regexpAddress); !ok {
				t.Fatalf("expected regexp address at end %T %v", end.Next, end.Next)
			}
		}
	}
	rest, _, err := bufReader.ReadRune()
	if err != nil {
		t.Fatalf("unexpected error reading rest %v", err)
	}
	if rest == 's' {
		t.Fatalf("unexpected rest %s", string(rest))
	}
}

func TestParseComposedAddres(t *testing.T) {
	bufReader := bufio.NewReader(strings.NewReader("1+2-d"))
	a, err := parseComposed(bufReader)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if addr, ok := a.(*composedAddress); !ok {
		t.Fatalf("unexpected address type, got %T", addr)
	}
	rest, _, err := bufReader.ReadRune()
	if err != nil {
		t.Fatalf("unexpected error reading rest %v", err)
	}
	if rest != 'd' {
		t.Fatalf("unexpected rest %s", string(rest))
	}

	bufReader = bufio.NewReader(strings.NewReader("/hello/1+2d"))
	a, err = parseComposed(bufReader)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if _, ok := a.(*composedAddress); !ok {
		t.Fatalf("unexpected address type, got %T", a)
	}
	rest, _, err = bufReader.ReadRune()
	if err != nil {
		t.Fatalf("unexpected error reading rest %v", err)
	}
	if rest != 'd' {
		t.Fatalf("unexpected rest %s", string(rest))
	}

	log.Print("Test 1")
	bufReader = bufio.NewReader(strings.NewReader("1\n"))
	a, err = parseComposed(bufReader)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if _, ok := a.(lineAddress); !ok {
		t.Fatalf("unexpected address type, got %T", a)
	}
}

func TestP(t *testing.T) {
	bufReader := bufio.NewReader(strings.NewReader("+1#2-d"))
	a, err := parseComposed(bufReader)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if addr, ok := a.(*composedAddress); !ok {
		t.Fatalf("unexpected address type, got %T", a)
	} else {
		if _, ok := addr.this.(*dotAddress); !ok {
			t.Fatalf("unexpected address type, got %T %v", addr.this, addr.this)
		}
		if addr2, ok := addr.Next.(*composedAddress); !ok {
			t.Fatalf("unexpected address type, got %T %v", addr.Next, addr.Next)
		} else {
			if _, ok := addr2.this.(lineAddress); !ok {
				t.Fatalf("unexpected address type, got %T %v", addr2.this, addr2.this)
			}
			if addr3, ok := addr2.Next.(*composedAddress); !ok {
				t.Fatalf("unexpected address type, got %T", addr2.Next)
			} else {
				if _, ok := addr3.this.(offsetAddress); !ok {
					t.Fatalf("unexpected address type, got %T %v", addr3.this, addr3.this)
				}
				if _, ok := addr3.Next.(lineAddress); !ok { // TODO should be the last line offset
					t.Fatalf("unexpected address type, got %T %v", addr3.Next, addr3.Next)
				}
			}
		}
	}
	rest, _, err := bufReader.ReadRune()
	if err != nil {
		t.Fatalf("unexpected error reading rest %v", err)
	}
	if rest != 'd' {
		t.Fatalf("unexpected rest %s", string(rest))
	}
}

func TestParseLineAddress(t *testing.T) {
	bufReader := bufio.NewReader(strings.NewReader("123"))
	addr, err := parseAddress(bufReader)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if lineAddr, ok := addr.(lineAddress); !ok {
		t.Fatalf("unexpected address type, got %T", addr)
	} else {
		if int(lineAddr) != 123 {
			t.Fatalf("unexpected offet address, got %v", lineAddr)
		}
	}

	addr, err = parseAddress(strings.NewReader("0123"))
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if lineAddr, ok := addr.(lineAddress); !ok {
		t.Fatalf("unexpected address type, got %T", addr)
	} else {
		if int(lineAddr) != 123 {
			t.Fatalf("unexpected offet address, got %v", lineAddr)
		}
	}

	bufReader = bufio.NewReader(strings.NewReader("123\n/abc/"))
	addr, err = parseAddress(bufReader)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if lineAddr, ok := addr.(lineAddress); !ok {
		t.Fatalf("unexpected address type, got %T", addr)
	} else {
		if int(lineAddr) != 123 {
			t.Fatalf("unexpected offet address, got %v", lineAddr)
		}
	}
	rest, _, err := bufReader.ReadRune()
	if err != nil {
		t.Fatalf("unexpected error reading rest %v", err)
	}
	if rest != '\n' {
		t.Fatalf("unexpected rest %v", rest)
	}
}

func TestParseOffsetAddress(t *testing.T) {
	addr, err := parseAddress(strings.NewReader("0"))
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if offsetAddr, ok := addr.(offsetAddress); !ok {
		t.Fatalf("unexpected address type, got %T", addr)
	} else {
		if int(offsetAddr) != 0 {
			t.Fatalf("unexpected offet address, got %v", addr)
		}
	}

	addr, err = parseAddress(strings.NewReader("#123"))
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if offsetAddr, ok := addr.(offsetAddress); !ok {
		t.Fatalf("unexpected address type, got %T", addr)
	} else {
		if int(offsetAddr) != 123 {
			t.Fatalf("unexpected offet address, got %v", addr)
		}
	}
}

func TestParseDotAddress(t *testing.T) {
	addr, err := parseAddress(strings.NewReader(""))
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if _, ok := addr.(*dotAddress); !ok {
		t.Fatalf("unexpected address type, got %T", addr)
	}

	bufReader := bufio.NewReader(strings.NewReader(".abc"))
	addr, err = parseAddress(bufReader)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if _, ok := addr.(*dotAddress); !ok {
		t.Fatalf("unexpected address type, got %T", addr)
	}
	rest, _, err := bufReader.ReadRune()
	if err != nil {
		t.Fatalf("unexpected error reading rest %v", err)
	}
	if rest != 'a' {
		t.Fatalf("unexpected rest %s", string(rest))
	}

	bufReader = bufio.NewReader(strings.NewReader("d"))
	addr, err = parseAddress(bufReader)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if _, ok := addr.(*dotAddress); !ok {
		t.Fatalf("unexpected address type, got %T", addr)
	}
	rest, _, err = bufReader.ReadRune()
	if err != nil {
		t.Fatalf("unexpected error reading rest %v", err)
	}
	if rest != 'd' {
		t.Fatalf("unexpected rest %s", string(rest))
	}
}

func TestParseRegexpAddress(t *testing.T) {
	addr, err := parseAddress(strings.NewReader("/findDot./"))
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if regexpAddr, ok := addr.(*regexpAddress); !ok {
		t.Fatalf("unexpected address type, got %T", addr)
	} else {
		if regexpAddr.regexp.String() != "findDot." {
			t.Fatalf("unexpected offet address, got '%s'", regexpAddr.regexp)
		}
		if regexpAddr.direction != forwards {
			t.Fatalf("unexpected direction, '%v'", regexpAddr.direction)
		}
	}

	addr, err = parseAddress(strings.NewReader("?findReverse.?"))
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if regexpAddr, ok := addr.(*regexpAddress); !ok {
		t.Fatalf("unexpected address type, got %T", addr)
	} else {
		if regexpAddr.regexp.String() != "findReverse." {
			t.Fatalf("unexpected offet address, got '%s'", regexpAddr.regexp)
		}
		if regexpAddr.direction != backwards {
			t.Fatalf("unexpected direction, '%v'", regexpAddr.direction)
		}
	}

	addr, err = parseAddress(strings.NewReader("?findReverse."))
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if regexpAddr, ok := addr.(*regexpAddress); !ok {
		t.Fatalf("unexpected address type, got %T", addr)
	} else {
		if regexpAddr.regexp.String() != "findReverse." {
			t.Fatalf("unexpected offet address, got '%s'", regexpAddr.regexp)
		}
		if regexpAddr.direction != backwards {
			t.Fatalf("unexpected direction, '%v'", regexpAddr.direction)
		}
	}

	bufReader := bufio.NewReader(strings.NewReader("?findReverse.?d"))
	addr, err = parseAddress(bufReader)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if regexpAddr, ok := addr.(*regexpAddress); !ok {
		t.Fatalf("unexpected address type, got %T", addr)
	} else {
		if regexpAddr.regexp.String() != "findReverse." {
			t.Fatalf("unexpected offet address, got '%s'", regexpAddr.regexp)
		}
		if regexpAddr.direction != backwards {
			t.Fatalf("unexpected direction, '%v'", regexpAddr.direction)
		}
	}
	rest, _, err := bufReader.ReadRune()
	if err != nil {
		t.Fatalf("unexpected error reading rest %v", err)
	}
	if rest != 'd' {
		t.Fatalf("unexpected rest %s", string(rest))
	}

	bufReader = bufio.NewReader(strings.NewReader("?findReverse.\nline 2"))
	addr, err = parseAddress(bufReader)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if regexpAddr, ok := addr.(*regexpAddress); !ok {
		t.Fatalf("unexpected address type, got %T", addr)
	} else {
		if regexpAddr.regexp.String() != "findReverse." {
			t.Fatalf("unexpected offet address, got '%s'", regexpAddr.regexp)
		}
		if regexpAddr.direction != backwards {
			t.Fatalf("unexpected direction, '%v'", regexpAddr.direction)
		}
	}
	rest, _, err = bufReader.ReadRune()
	if err != nil {
		t.Fatalf("unexpected error reading rest %v", err)
	}
	if rest != 'l' {
		t.Fatalf("unexpected rest %s", string(rest))
	}
}
