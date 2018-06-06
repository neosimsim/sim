package main

import (
	"testing"
)

func (file *File) assertBuffer(t *testing.T, expected string) {
	if file.Buffer != expected {
		t.Errorf("expected buffer '%s' got '%s'", expected, file.Buffer)
	}
}


func TestAppend(t *testing.T) {
	file := File{Buffer: "Hello World!\n", Adress: Range{2, 6}}
	newTest := "Big "
	file.Append(newTest)

	file.assertBuffer(t, "Hello Big World!\n");
	if file.Adress.Start != 6 || file.Adress.End != 6+len(newTest) {
		t.Error("range should have been set")
	}
}

func BenchmarkAppend(b *testing.B) {
	file := File{Buffer: "Hello World!\n", Adress: Range{2, 6}}
	for i := 0; i < b.N; i++ {
		file.Append("Big ")
	}
}

func TestChange(t *testing.T) {
	file := File{Buffer: "Hello World!", Adress: Range{6, 11}}
	file.Change("Moon")

	file.assertBuffer(t, "Hello Moon!")
	if file.Adress.Start != 6 || file.Adress.End != 10 {
		t.Error("range should have been updated")
	}
}

func BenchmarkChange(b *testing.B) {
	file := File{Buffer: "Hello World!\n", Adress: Range{6, 11}}
	for i := 0; i < b.N; i++ {
		file.Append("Moon")
	}
}

func TestDeleteDot(t *testing.T) {
	file := File{Buffer: "Hello World!", Adress: Range{5, 11}}
	file.DeleteDot()

	file.assertBuffer(t, "Hello!")
	if file.Adress.Start != 5 || file.Adress.End != 5 {
		t.Error("range should have been set")
	}
}

func TestCopyDot(t *testing.T) {
	file := File{Buffer: "Hello World!", Adress: Range{5, 11}}
	file.CopyDot(5)

	file.assertBuffer(t, "Hello World World!")
	if file.Adress.Start != 6 || file.Adress.End != 10 {
		t.Error("range should have been set")
	}
}

func BenchmarkCopyDot(b *testing.B) {
	file := File{Buffer: "Hello World!", Adress: Range{6, 11}}
	for i := 0; i < b.N; i++ {
		file.Append("Moon")
	}
}

func TestMoveDot(t *testing.T) {
	file := File{Buffer: "Hello World!", Adress: Range{6, 11}}
	file.MoveDot(2)

	file.assertBuffer(t, "HeWorldlo !")
	if file.Adress.Start != 6 || file.Adress.End != 10 {
		t.Error("range should have been set")
	}
}

func BenchmarkMoveDot(b *testing.B) {
	file := File{Buffer: "Hello World!\n", Adress: Range{6, 11}}
	for i := 0; i < b.N; i++ {
		file.Append("Moon")
	}
}

func TestDot(t *testing.T) {
	file := File{Buffer: "Hello World!\n", Adress: Range{3, 7}}
	if file.Dot() != "lo W" {
		t.Errorf("expected '%s' got '%s'", "lo W", file.Dot())
	}
}

// Thist test shall assure the the zero value of file.File is useful.
func TestZeroValueOfFile(t *testing.T) {
	var file File
	if file.Buffer != "" {
		t.Error("Zero file should have the empty string as buffer")
	}
	if file.Adress.Start !=0 && file.Adress.End != 0 {
		t.Error("Zero file should have the empty address at index 0")
	}
}
