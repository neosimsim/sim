package parse

import (
	"regexp"

	"github.com/neosimsim/sim"
)

// AddressDesc is an interface to provide functions to scan a file for an
// Address.
type AddressDesc interface {
	// Scan for the Address in File described by AddrresDesc after Address start.
	Scan(f *sim.File, start sim.Address) (sim.Address, error)
}

type RegexpOffset regexp.Regexp
func (o RegexpOffset) Scan(f *sim.File, start sim.Address) (sim.Address, error) {
	return 0, nil
}

type LineOffset uint
func (o LineOffset) Scan(f *sim.File, start sim.Address) (sim.Address, error) {
	return 0, nil
}

type CharOffset uint
func (o CharOffset) Scan(f *sim.File, start sim.Address) (sim.Address, error) {
	return 0, nil
}
