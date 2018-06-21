// Copyright © 2018, Alexander Ben Nasrallah <me@abn.sh>
// Use of this source code is governed by a BSD 3-clause
// style license that can be found in the LICENSE file.

package main

import (
	"strings"
)

type Address struct {
	Start int
	End   int
}

type InvalidAddressError struct {
	Addr Address
}

func NewInvalidAddressError(addr Address) *InvalidAddressError {
	return &InvalidAddressError{Addr: addr}
}

func (*InvalidAddressError) Error() string {
	return ""
}

// A  File is an in memory copy of an external file, e.g. stored on disk.
type File struct {
	FileName string
	clean    bool
	buffer   string
	dot      Address
	mark     Address
}

func (f *File) SetText(text string, addr Address) (Address, error) {
	if err := f.validateAddress(addr); err != nil {
		return Address{}, err
	}
	f.buffer = strings.Join([]string{f.buffer[:addr.Start], text, f.buffer[addr.End:]}, "")
	return Address{Start: addr.Start, End: addr.Start + len(text)}, nil
}

func (f *File) GetText(addr Address) (string, error) {
	if err := f.validateAddress(addr); err != nil {
		return "", err
	}
	return f.buffer[addr.Start:addr.End], nil
}

func (f *File) validateAddress(addr Address) error {
	if addr.Start < 0 || addr.End < 0 || addr.Start > addr.End || addr.Start > len(f.buffer) || addr.End > len(f.buffer) {
		return &InvalidAddressError{addr}
	}
	return nil
}
