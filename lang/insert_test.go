package lang

import (
	"bufio"
	"strings"
	"testing"
)

func TestReadInsert(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(`/new text to insert/`))
	insert, err := ParseInsert(reader)

	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	if insert == nil {
		t.Fatal("expected Insert got nil")
	}

	if insert.Text != `new text to insert` {
		t.Errorf("unexpected text to be inserted %v", insert.Text)
	}
}

func TestReadInsertBlock(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(`
new text

.	
	to insert
.
`))

	insert, err := ParseInsert(reader)

	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	if insert == nil {
		t.Fatal("expected Insert got nil")
	}

	if insert.Text != "new text\n\n.\t\n\tto insert\n" {
		t.Errorf("unexpected text to be inserted %v", insert.Text)
	}
}
