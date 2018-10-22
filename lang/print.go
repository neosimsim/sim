package lang

import (
	"github.com/neosimsim/sim/file"
)

type Print struct {
}

func (cmd *Print) Process(f file.File, addr Address) ([]FileModification, error) {
	return nil, nil
}
