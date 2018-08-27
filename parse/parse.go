package parse

import (
	"bufio"

	"github.com/neosimsim/sim"
)

type Command struct {
	Address AddressDesc
	Instruction Instruction
}

func (cmd *Command) Process(file sim.File) ([]FileModification, error) {
	return cmd.Instruction.Process(file, cmd.Address)
}

type FileModification struct {
	Start int
	End int
	Text string
}

func Parse(r *bufio.Reader) (*Command, error) {
	return nil, nil
}

func ParseAddress(r *bufio.Reader) (AddressDesc, error) {
	return nil, nil
}
