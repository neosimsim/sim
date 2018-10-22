package lang

import (
	"io"

	"github.com/neosimsim/sim/file"
)

type Command struct {
	a1          Address
	Instruction Instruction
}

type Instruction interface {
	Process(f file.File, addr Address) ([]FileModification, error)
}

func (cmd *Command) Process(file file.File) ([]FileModification, error) {
	return cmd.Instruction.Process(file, cmd.a1)
}

type FileModification struct {
	Start int
	End   int
	Text  string
}

func ParseCommand(r io.Reader) (*Command, error) {
	addr, err := ParseAddress(r)
	if err != nil {
		return nil, err
	}
	instr, err := ParseInstruction(r)
	if err != nil {
		return nil, err
	}
	return &Command{a1: addr, Instruction: instr}, nil
}

func ParseInstruction(r io.Reader) (Instruction, error) {
	reader := NewLangReader(r)
	cmdName, _ := reader.ReadWord()
	switch cmdName {
	case "a":
	}
	return nil, nil
}
