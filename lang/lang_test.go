package lang

import (
	"bufio"
	"strings"
	"testing"
)

func TestParseCommand(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(`2a/new text to append/`))
	command, err := ParseCommand(reader)

	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	if command == nil {
		t.Fatal("expected command got nil")
	}

	if addr, ok := command.a1.(lineAddress); !ok {
		t.Errorf("expected line address, got %T %v", command.a1, command.a1)
	} else if addr != 2 {
		t.Errorf("expected line address %d", addr)
	}
	if instruction, ok := command.Instruction.(*Append); !ok {
		t.Errorf("expected Append, got %T %v", command.Instruction, command.Instruction)
	} else if instruction.Text != "new text to append" {
		t.Errorf("expected text to append '%s'", instruction.Text)
	}
}
