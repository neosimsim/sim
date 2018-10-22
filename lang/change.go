package lang

import (
	"io"

	"github.com/neosimsim/sim/file"
)

type Change struct {
	Text string
}

func ParseChange(r io.Reader) (*Change, error) {
	reader := NewLangReader(r)
	if text, err := reader.ReadLineOrBlock(); err != nil {
		return nil, err
	} else {
		return &Change{Text: text}, nil
	}
}

func (append *Change) Process(f file.File, addr Address) ([]FileModification, error) {
	return nil, nil
}
