package lang

import (
	"bufio"
	"io"
	"strings"
	"testing"
)

func TestReadLineToDelim(t *testing.T) {
	bufReader := bufio.NewReader(strings.NewReader(`new\/text to append/a`))
	reader := NewLangReader(bufReader)
	text, err := reader.ReadLineTo('/')

	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	if text != `new/text to append` {
		t.Errorf("unexpected text '%v'", text)
	}
	rest, _, err := bufReader.ReadRune()
	if err != nil {
		t.Errorf("unexpected error reading rest %v", err)
	}
	if rest != 'a' {
		t.Errorf("unexpected rest %q", rest)
	}
}

func TestReadLineToEol(t *testing.T) {
	reader := NewLangReader(strings.NewReader("new text to append\n"))
	text, err := reader.ReadLineTo('/')

	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	if text != `new text to append` {
		t.Errorf("unexpected text '%v'", text)
	}
}

func TestReadLineToEof(t *testing.T) {
	reader := NewLangReader(strings.NewReader("new text to append"))
	text, err := reader.ReadLineTo('/')

	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	if text != `new text to append` {
		t.Errorf("unexpected text '%v'", text)
	}
}

func TestBlock(t *testing.T) {
	reader := NewLangReader(strings.NewReader("new text\n.to\n.\nappend"))
	text, err := reader.ReadBlock()

	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	if text != "new text\n.to\n" {
		t.Errorf("unexpected text '%v'", text)
	}
}

func TestBlockAtEOF(t *testing.T) {
	reader := NewLangReader(strings.NewReader("new text\n.to\n."))
	_, err := reader.ReadBlock()

	if err != io.EOF {
		t.Errorf("unexpected error %v", err)
	}
}

func TestReadDelim(t *testing.T) {
	reader := NewLangReader(strings.NewReader(`/new text to append/`))

	delim, err := reader.ReadDelim()

	if err != nil {
		t.Errorf("unexpected error '%v'", err)
	}

	if delim != '/' {
		t.Errorf("unexpected delim %q", delim)
	}
}

func TestReadDelimSkipWhitespace(t *testing.T) {
	reader := NewLangReader(strings.NewReader(`      /new text to append/`))

	delim, err := reader.ReadDelim()

	if err != nil {
		t.Errorf("unexpected error '%v'", err)
	}

	if delim != '/' {
		t.Errorf("unexpected delim %q", delim)
	}
}

func TestReadInvalidDelim(t *testing.T) {
	reader := NewLangReader(strings.NewReader("hallo welt\n."))

	delim, err := reader.ReadDelim()

	if err != InvalidDelim {
		t.Errorf("unexpected error %v", err)
	}

	if delim != 'h' {
		t.Errorf("unexpected delim %q", delim)
	}
}

func TestReadEolDelim(t *testing.T) {
	reader := NewLangReader(strings.NewReader("\nhallo welt\n."))

	delim, err := reader.ReadDelim()

	if err != EOL {
		t.Errorf("unexpected error '%v'", err)
	}

	if delim != ' ' {
		t.Errorf("unexpected delim %q", delim)
	}
}

func TestReadNumber(t *testing.T) {
	bufReader := bufio.NewReader(strings.NewReader("232abc"))
	reader := NewLangReader(bufReader)
	number, err := reader.ReadNumber()
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if number != 232 {
		t.Errorf("unexpected number %d", number)
	}
	rest, _, err := bufReader.ReadRune()
	if err != nil {
		t.Fatalf("unexpected error reading rest %v", err)
	}
	if rest != 'a' {
		t.Errorf("expected reader to have stopped after number: %v", string(rest))
	}

	reader = NewLangReader(strings.NewReader(""))
	number, err = reader.ReadNumber()
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if number != 0 {
		t.Errorf("unexpected number %d", number)
	}

	bufReader = bufio.NewReader(strings.NewReader("abc"))
	reader = NewLangReader(bufReader)
	number, err = reader.ReadNumber()
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if number != 0 {
		t.Errorf("unexpected number %d", number)
	}
	rest, _, err = bufReader.ReadRune()
	if err != nil {
		t.Fatalf("unexpected error reading rest %v", err)
	}
	if rest != 'a' {
		t.Errorf("expected reader to have stopped after number: %v", string(rest))
	}
}

func TestReadWord(t *testing.T) {
	bufReader := bufio.NewReader(strings.NewReader("cp/xxx"))
	reader := NewLangReader(bufReader)
	word, err := reader.ReadWord()
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if word != "cp" {
		t.Errorf("unexpected word '%s'", word)
	}
	rest, _, err := bufReader.ReadRune()
	if err != nil {
		t.Fatalf("unexpected error reading rest %v", err)
	}
	if rest != '/' {
		t.Errorf("expected reader to have stopped after word: %v", string(rest))
	}

	reader = NewLangReader(strings.NewReader(""))
	word, err = reader.ReadWord()
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if word != "" {
		t.Errorf("unexpected word '%s'", word)
	}

	bufReader = bufio.NewReader(strings.NewReader("123"))
	reader = NewLangReader(bufReader)
	word, err = reader.ReadWord()
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if word != "" {
		t.Errorf("unexpected word '%s'", word)
	}
	rest, _, err = bufReader.ReadRune()
	if err != nil {
		t.Fatalf("unexpected error reading rest %v", err)
	}
	if rest != '1' {
		t.Errorf("expected reader to have stopped after word: '%s'", string(rest))
	}
}
