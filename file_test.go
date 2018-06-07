package main

import (
	"testing"
)

func (file *File) assertBuffer(expected string, t *testing.T) {
	if file.buffer != expected {
		t.Errorf("expected buffer '%s' got '%s'", expected, file.buffer)
	}
}

func (file *File) assertDot(start, end int, t *testing.T) {
	if file.dot.Start != start || file.dot.End != end {
		t.Errorf("expected dot #%d,#%d got #%d,#%d", start, end, file.dot.Start, file.dot.End)
	}
}

func TestAppend(t *testing.T) {
	file := File{buffer: "Hello World!\n"}
	newTest := "Big "
	file.Append(newTest, Address{2, 6})

	file.assertBuffer("Hello Big World!\n", t)
	file.assertDot(6, 6+len(newTest), t)
}

func TestAppendOnInvalidRange(t *testing.T) {
	file := File{buffer: "Hello World!"}
	err := file.Append("Big ", Address{0, 6})
	if err == nil {
		t.Fatalf("expected InvalidAddresError got %v", err)
	}
	addrErr := err.(*InvalidAddressError)
	if addrErr == nil {
		t.Error("expected an error")
	}
}

func TestChange(t *testing.T) {
	file := File{buffer: "Hello World!"}
	file.Change("Moon", Address{6, 11})

	file.assertBuffer("Hello Moon!", t)
	file.assertDot(6, 10, t)
}

func TestDelete(t *testing.T) {
	file := File{buffer: "Hello World!"}
	file.Delete(Address{5, 11})

	file.assertBuffer("Hello!", t)
	file.assertDot(5, 5, t)
}

func TestCopy(t *testing.T) {
	file := File{buffer: "Hello World!"}
	file.Copy(Address{Start: 5, End: 10}, Address{Start: 3, End: 5})

	file.assertBuffer("Hello World World!", t)
	file.assertDot(6, 10, t)
}


func TestMove(t *testing.T) {
	file := File{buffer: "Hello World!"}
	file.Move(Address{Start: 2, End: 2}, Address{Start: 6, End: 10})

	file.assertBuffer("HeWorldlo !", t)
	file.assertDot(6, 10, t)
}

func TestPrint(t *testing.T) {
	file := File{buffer: "Hello World!\n"}
	result, err := file.Print(Address{Start: 3, End: 6})
	if result != "lo W" {
		t.Errorf("expected '%s' got '%s'", "lo W", result)
	}
	if err != nil {
		t.Error("unexpected err", err)
	}
	file.assertDot(3, 6, t)
}

// Thist test shall assure the the zero value of file.File is useful.
func TestZeroValueOfFile(t *testing.T) {
	var file File
	if file.buffer != "" {
		t.Error("Zero file should have the empty string as buffer")
	}
	if file.dot.Start != 0 && file.dot.End != 0 {
		t.Error("Zero file should have the empty address at index 0")
	}
}

func TestApplyMark(t *testing.T) {
	t.Error("Not implemented")
}
