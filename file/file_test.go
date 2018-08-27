package file

import (
	"testing"
)

func TestInsert(t *testing.T) {
	file := &File{buffer: "Hello World!", dot: Range{Start: 7, End: 11}}
	addr, err := file.Insert(" new test content ", 3)

	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if file.buffer != "Hel new test content lo World!" {
		t.Errorf("expected buffer '%s' got '%s'", "Hel new test content rld!", file.buffer)
	}
	if file.dot.Start != 7 || file.dot.End != 11 {
		t.Errorf("expected dot to be unchanged #%d,#%d got #%d,#%d", 7, 11, file.dot.Start, file.dot.End)
	}
	if addr.Start != 3 || addr.End != 21 {
		t.Errorf("expected returned address #%d,#%d got #%d,#%d", 3, 21, addr.Start, addr.End)
	}
}

func TestInsertOnInvalidAddress(t *testing.T) {
	file := &File{buffer: "Hello World!", dot: Range{Start: 7, End: 11}}

	_, err := file.Insert(" new test content ", 30)
	if _, ok := err.(*InvalidAddressError); !ok {
		t.Errorf("expected error of type *InvalidAddressError for Start > len got %T", err)
	}
}

func TestGetText(t *testing.T) {
	file := &File{buffer: "Hello World!", dot: Range{Start: 7, End: 11}}
	text, err := file.GetText(Range{Start: 5, End: 9})

	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if text != " Wor" {
		t.Errorf("expected text '%s' got '%s'", " Wor", text)
	}
	if file.dot.Start != 7 || file.dot.End != 11 {
		t.Errorf("expected dot to be unchanged #%d,#%d got #%d,#%d", 7, 11, file.dot.Start, file.dot.End)
	}
}

func TestGetTextOnInvalidRange(t *testing.T) {
	file := &File{buffer: "Hello World!", dot: Range{Start: 7, End: 11}}

	_, err := file.GetText(Range{Start: 30, End: 40})
	if _, ok := err.(*InvalidAddressError); !ok {
		t.Errorf("expected error of type *InvalidAddressError for Start > len got %T", err)
	}

	_, err = file.GetText(Range{Start: 1, End: 40})
	if _, ok := err.(*InvalidAddressError); !ok {
		t.Errorf("expected error of type *InvalidAddressError for End > len got %T", err)
	}

	_, err = file.GetText(Range{Start: 8, End: 4})
	if _, ok := err.(*InvalidRangeError); !ok {
		t.Errorf("expected error of type *InvalidRangeError for Start > End got %T", err)
	}
}

func TestDelete(t *testing.T) {
	file := &File{buffer: "Hello World!", dot: Range{Start: 7, End: 11}}

	length, err := file.Delete(Range{Start: 4, End: 9})
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	if file.buffer != "Hellld!" {
		t.Errorf("expected buffer '%s' got '%s'", "Hellld!", file.buffer)
	}

	if length != 5 {
		t.Errorf("expected to delete %d runes got %d", 5, length)
	}
}

func TestDeleteOnInvalidRange(t *testing.T) {
	file := &File{buffer: "Hello World!", dot: Range{Start: 7, End: 11}}

	_, err := file.Delete(Range{Start: 30, End: 40})
	if _, ok := err.(*InvalidAddressError); !ok {
		t.Errorf("expected error of type *InvalidAddressError for Start > len got %T", err)
	}

	_, err = file.Delete(Range{Start: 1, End: 40})
	if _, ok := err.(*InvalidAddressError); !ok {
		t.Errorf("expected error of type *InvalidAddressError for End > len got %T", err)
	}

	_, err = file.Delete(Range{Start: 8, End: 4})
	if _, ok := err.(*InvalidRangeError); !ok {
		t.Errorf("expected error of type *InvalidRangeError for Start > End got %T", err)
	}
}
