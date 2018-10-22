package lang

import (
	"github.com/neosimsim/sim/file"
)

type Delete struct {
}

func (append *Delete) Process(f file.File, addr Address) ([]FileModification, error) {
	return nil, nil
}
