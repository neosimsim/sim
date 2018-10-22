package lang

import (
	"io"

	"github.com/neosimsim/sim/file"
)

type Insert struct {
	Text string
}

func ParseInsert(r io.Reader) (*Insert, error) {
	reader := NewLangReader(r)
	if text, err := reader.ReadLineOrBlock(); err != nil {
		return nil, err
	} else {
		return &Insert{Text: text}, nil
	}
}

func (cmd *Insert) Process(f file.File, addr Address) ([]FileModification, error) {
	return nil, nil
}
