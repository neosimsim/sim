package lang

import (
	"io"

	"github.com/neosimsim/sim/file"
)

type Append struct {
	Text string
}

func ParseAppend(r io.Reader) (*Append, error) {
	reader := NewLangReader(r)
	if text, err := reader.ReadLineOrBlock(); err != nil {
		return nil, err
	} else {
		return &Append{Text: text}, nil
	}
}

func (append *Append) Process(f file.File, addr Address) ([]FileModification, error) {
	return nil, nil
}
