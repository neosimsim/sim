package parse

import (
	"regexp"

	"github.com/neosimsim/sim/file"
)

// AddressDesc is an interface to provide functions to scan a file for an
// Address.
type AddressDesc interface {
	// Scan for the Address in File described by AddrresDesc after Address start.
	Scan(f *file.File, start file.Address) (file.Address, error)
}

type RegexpOffset regexp.Regexp
func (o RegexpOffset) Scan(f *file.File, start file.Address) (file.Address, error) {
	return 0, nil
}

type LineOffset uint
func (o LineOffset) Scan(f *file.File, start file.Address) (file.Address, error) {
	return 0, nil
}

type CharOffset uint
func (o CharOffset) Scan(f *file.File, start file.Address) (file.Address, error) {
	return 0, nil
}
