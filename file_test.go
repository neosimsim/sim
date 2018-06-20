package main

import (
	"testing"
)

func TestAppend(t *testing.T) {
	file := NewTestFile()
	assertFileChange(t)(file, func(a Address) error { return file.Append(" added content", a) }, "Hello Wo added content rld!")
}

func TestChange(t *testing.T) {
	file := NewTestFile()
	assertFileChange(t)(file, func(a Address) error { return file.Change("changed content", a) }, "Hellchanged contentrld!")
}

func TestCopy(t *testing.T) {
	file := NewTestFile()
	assertFileChange(t)(file, func(a Address) error { return file.Copy(Address{4, 6}, a) }, "Hello lo WoWorld!")
}

func TestDelete(t *testing.T) {
	file := NewTestFile()
	assertFileChange(t)(file, file.Delete, "Helrld!")
}

func TestInsert(t *testing.T) {
	file := NewTestFile()
	assertFileChange(t)(file, func(a Address) error { return file.Insert(" inserted content ", a) }, "Hel inserted content lo World!")
}

func TestMove(t *testing.T) {
	file := NewTestFile()
	assertFileChange(t)(file, func(a Address) error { return file.Copy(Address{4, 9}, a) }, "Helrlo Wold!")

	assertFileChange(t)(file, func(a Address) error { return file.Copy(Address{4, 6}, a) }, "Hel lo World!")
	t.Error("Address overlapping")
}

func TestPipe(t *testing.T) {
	file := NewTestFile()
	assertFileChange(t)(file, func(a Address) error { return file.Pipe("", a) }, "Hel loWorld!")
	t.Error("Not implemented")
}

func TestPipeIn(t *testing.T) {
	file := NewTestFile()
	assertFileChange(t)(file, func(a Address) error { return file.PipeOut("", a) }, "Hel loWorld!")
	t.Error("Not implemented")
}

func TestPipeOut(t *testing.T) {
	file := NewTestFile()
	assertFileChange(t)(file, func(a Address) error { return file.PipeIn("", a) }, "Hel loWorld!")
	t.Error("Not implemented")
}

func TestPrint(t *testing.T) {
	file := NewTestFile()
	assertFileChange(t)(file, func(a Address) error { file.Print(a); return nil }, "Hel loWorld!")
	t.Error("Maybe Print should be called Dot")
}

func TestSetMark(t *testing.T) {
	file := NewTestFile()
	assertFileChange(t)(file, file.SetMark, "Hel loWorld!")
}

func TestSubstitute(t *testing.T) {
	file := NewTestFile()
	assertFileChange(t)(file, func(a Address) error { return file.Substitute("..$", "xyz", a) }, "Hello xyzrld!")
}

func NewTestFile() *File {
	return &File{buffer: "Hello World!", dot: Address{Start: 9, End: 11}}
}

func assertFileChange(t *testing.T) func(*File, func(Address) error, string) {
	return func(file *File, f func(Address) error, expectedBuffer string) {
		err := f(Address{Start: 3, End: 8})

		if err != nil {
			t.Errorf("Unexpected error %v", err)
		}

		if file.buffer != expectedBuffer {
			t.Errorf("expected buffer '%s' got '%s'", expectedBuffer, file.buffer)
		}
		if file.dot.Start != 9 || file.dot.End != 11 {
			t.Errorf("expected dot to be unchanged #%d,#%d got #%d,#%d", 9, 11, file.dot.Start, file.dot.End)
		}

		err = f(Address{Start: -1, End: 5})
		if _, ok := err.(*InvalidAddressError); !ok {
			t.Errorf("expected error of type *InvalidAddressError for Start < 0 got %T", err)
		}

		err = f(Address{Start: 30, End: 40})
		if _, ok := err.(*InvalidAddressError); !ok {
			t.Errorf("expected error of type *InvalidAddressError for Start > len got %T", err)
		}

		err = f(Address{Start: 1, End: 40})
		if _, ok := err.(*InvalidAddressError); !ok {
			t.Errorf("expected error of type *InvalidAddressError for End > len got %T", err)
		}

		err = f(Address{Start: 8, End: 4})
		if _, ok := err.(*InvalidAddressError); !ok {
			t.Errorf("expected error of type *InvalidAddressError for Start > End got %T", err)
		}
	}
}
