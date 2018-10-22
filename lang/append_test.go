package lang

import (
	"bufio"
	"strings"
	"testing"
)

func TestReadAppend(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(`/new text to append/`))
	append, err := ParseAppend(reader)

	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	if append == nil {
		t.Fatal("expected Append got nil")
	}

	if append.Text != `new text to append` {
		t.Errorf("unexpected text to be appended %v", append.Text)
	}
}

func TestReadAppendBlock(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(`
new text

.	
	to append
.
`))

	append, err := ParseAppend(reader)

	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	if append == nil {
		t.Fatal("expected Append got nil")
	}

	if append.Text != "new text\n\n.\t\n\tto append\n" {
		t.Errorf("unexpected text to be appended %v", append.Text)
	}
}
