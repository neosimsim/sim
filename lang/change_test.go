package lang

import (
	"bufio"
	"strings"
	"testing"
)

func TestReadChange(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(`/new text to change/`))
	change, err := ParseChange(reader)

	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	if change == nil {
		t.Fatal("expected Change got nil")
	}

	if change.Text != `new text to change` {
		t.Errorf("unexpected text to be changeed %v", change.Text)
	}
}

func TestReadChangeBlock(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(`
new text

.	
	to change
.
`))

	change, err := ParseChange(reader)

	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	if change == nil {
		t.Fatal("expected Change got nil")
	}

	if change.Text != "new text\n\n.\t\n\tto change\n" {
		t.Errorf("unexpected text to be changeed %v", change.Text)
	}
}
