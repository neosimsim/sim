package lang

import (
	"testing"
	"strings"
	"bufio"
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

