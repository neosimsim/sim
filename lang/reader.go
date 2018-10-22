package lang

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"unicode"
)

var (
	EOL          = errors.New("reached EOL")
	InvalidDelim = errors.New("invalid delim")
)

// LangReader provides read function to read components of the
// sam command language.
type LangReader struct {
	bufReader *bufio.Reader
}

func NewLangReader(r io.Reader) *LangReader {
	return &LangReader{bufReader: bufio.NewReader(r)}
}

// Read until end of or delim is read.
// Returns the read string without delim or '\n'
// The delimiter can be escaped by '\'
func (reader *LangReader) ReadLineTo(delim rune) (string, error) {
	var text string
	for {
		c, _, err := reader.bufReader.ReadRune()
		if err == io.EOF {
			return text, nil
		}
		if err != nil {
			return "", err
		}
		if c == delim || c == '\n' {
			return text, nil
		}
		if c == '\\' {
			c, _, err = reader.bufReader.ReadRune()
		}
		text = fmt.Sprint(text, string(c))
	}
	return text, nil
}

func (reader *LangReader) ReadBlock() (string, error) {
	var text string
	for {
		line, err := reader.bufReader.ReadString('\n')
		if err != nil {
			return "", err
		}
		if line == ".\n" {
			return text, nil
		}
		text = fmt.Sprint(text, line)
	}
	return text, nil
}

// Reads until the next non space non letter rune.
// Returns the read rune and InvalidDelim if another rune has been read.
// Returns ' ', EOL if '\n' is read.
func (reader *LangReader) ReadDelim() (rune, error) {
	delim := ' '
	var err error
	for unicode.IsSpace(delim) && delim != '\n' {
		delim, _, err = reader.bufReader.ReadRune()
	}
	if err != nil {
		return ' ', err
	}
	if delim == '\n' {
		return ' ', EOL
	}
	if unicode.IsLetter(delim) || unicode.IsDigit(delim) {
		return delim, InvalidDelim
	}
	return delim, nil
}

func (reader *LangReader) ReadWord() (word string, err error) {
	for true {
		var nextRune rune
		nextRune, _, err = reader.bufReader.ReadRune()
		if err == io.EOF {
			return word, nil
		}
		if !unicode.IsLetter(nextRune) {
			err = reader.bufReader.UnreadRune()
			return
		}
		word += string(nextRune)
	}
	if err != nil && err != io.EOF {
		return "", err
	}
	return
}

func (reader *LangReader) ReadNumber() (number int, err error) {
	for true {
		var nextRune rune
		nextRune, _, err = reader.bufReader.ReadRune()
		if err == io.EOF {
			return number, nil
		}
		if !unicode.IsDigit(nextRune) {
			err = reader.bufReader.UnreadRune()
			return
		}
		digit, _ := strconv.Atoi(string(nextRune))
		number = number*10 + digit
	}
	if err != nil && err != io.EOF {
		return 0, err
	}
	return
}
