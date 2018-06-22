// Copyright © 2018, Alexander Ben Nasrallah <me@abn.sh>
// Use of this source code is governed by a BSD 3-clause
// style license that can be found in the LICENSE file.

package main

import (
	"strings"
)

type Address uint
type Range struct {
	Start Address 
	End   Address
}

type InvalidAddressError struct {
	Address Address
}

func NewInvalidAddressError(a Address) *InvalidAddressError {
	return &InvalidAddressError{Address: a}
}

func (*InvalidAddressError) Error() string {
	return ""
}

type InvalidRangeError struct {
	Range Range
}

func NewInvalidRangeError(r Range) *InvalidRangeError {
	return &InvalidRangeError{Range: r}
}

func (*InvalidRangeError) Error() string {
	return ""
}

// A  File is an in memory copy of an external file, e.g. stored on disk.
type File struct {
	FileName string
	clean    bool
	buffer   string
	dot      Range
	mark     Range
}

func (f *File) setText(text string, addr Range) (Range, error) {
	if err := f.validateRange(addr); err != nil {
		return Range{}, err
	}
	f.buffer = strings.Join([]string{f.buffer[:addr.Start], text, f.buffer[addr.End:]}, "")
	return Range{Start: addr.Start, End: addr.Start + Address(len(text))}, nil
}

// Insert the given text to the address of the file.
// Returning the Range of the added text.
func (f *File) Insert(text string, addr Address) (Range, error) {
	if err := f.validateAddress(addr); err != nil {
		return Range{}, err
	}
	return f.setText(text, Range{Start: addr, End: addr})
}

// Delete the range from the file.
// Return the length of the deleted range.
func (f *File) Delete(r Range) (Address, error) {
	if err := f.validateRange(r); err != nil {
		return 0, err
	}
	f.setText("", r)
	return r.End - r.Start, nil
}

func (f *File) GetText(addr Range) (string, error) {
	if err := f.validateRange(addr); err != nil {
		return "", err
	}
	return f.buffer[addr.Start:addr.End], nil
}

func (f *File) validateRange(r Range) error {
	if err := f.validateAddress(r.Start); err != nil {
		return err
	}
	if err := f.validateAddress(r.End); err != nil {
		return err
	}
	if r.Start > r.End {
		return &InvalidRangeError{r}
	}
	return nil
}

func (f *File) validateAddress(addr Address) error {
	if addr > Address(len(f.buffer)) {
		return &InvalidAddressError{addr}
	}
	return nil
}
