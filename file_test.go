package main

import (
	"testing"
)

func TestSetText(t *testing.T) {
	file := &File{buffer: "Hello World!", dot: Address{Start: 7, End: 11}}
	addr, err := file.SetText(" new test content ", Address{Start: 3, End: 8})

	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if file.buffer != "Hel new test content rld!" {
		t.Errorf("expected buffer '%s' got '%s'", "Hel new test content rld!", file.buffer)
	}
	if file.dot.Start != 7 || file.dot.End != 11 {
		t.Errorf("expected dot to be unchanged #%d,#%d got #%d,#%d", 7, 11, file.dot.Start, file.dot.End)
	}
	if addr.Start != 3 || addr.End != 21 {
		t.Errorf("expected returned address #%d,#%d got #%d,#%d", 3, 21, addr.Start, addr.End)
	}
}

func TestSetTextInEmptyAddress(t *testing.T) {
	file := &File{buffer: "Hello World!", dot: Address{Start: 7, End: 11}}
	addr, err := file.SetText(" new test content ", Address{Start: 5, End: 5})

	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if file.buffer != "Hello new test content  World!" {
		t.Errorf("expected buffer '%s' got '%s'", "Hel new test content rld!", file.buffer)
	}
	if file.dot.Start != 7 || file.dot.End != 11 {
		t.Errorf("expected dot to be unchanged #%d,#%d got #%d,#%d", 7, 11, file.dot.Start, file.dot.End)
	}
	if addr.Start != 5 || addr.End != 23 {
		t.Errorf("expected returned address #%d,#%d got #%d,#%d", 5, 23, addr.Start, addr.End)
	}
}

func TestErrorOnInvalidAddress(t *testing.T) {
	file := &File{buffer: "Hello World!", dot: Address{Start: 7, End: 11}}

	_, err := file.SetText(" new test content ", Address{Start: -1, End: 8})
	if _, ok := err.(*InvalidAddressError); !ok {
		t.Errorf("expected error of type *InvalidAddressError for Start < 0 got %T", err)
	}

	_, err = file.SetText(" new test content ", Address{Start: 30, End: 40})
	if _, ok := err.(*InvalidAddressError); !ok {
		t.Errorf("expected error of type *InvalidAddressError for Start > len got %T", err)
	}

	_, err = file.SetText(" new test content ", Address{Start: 1, End: 40})
	if _, ok := err.(*InvalidAddressError); !ok {
		t.Errorf("expected error of type *InvalidAddressError for End > len got %T", err)
	}

	_, err = file.SetText(" new test content ", Address{Start: 8, End: 4})
	if _, ok := err.(*InvalidAddressError); !ok {
		t.Errorf("expected error of type *InvalidAddressError for Start > End got %T", err)
	}
}
