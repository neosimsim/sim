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
	delim, err := reader.ReadDelim()
	var text string
	if err != nil {
		if err == EOL {
			text, err = reader.ReadBlock()
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		text, err = reader.ReadLineTo(delim)
		if err != nil {
			return nil, err
		}
	}
	return &Append{Text: text}, nil
}

func (append *Append) Process(f file.File, addr Address) ([]FileModification, error) {
	return nil, nil
}
