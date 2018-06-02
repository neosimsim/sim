package main

import (
	"testing"
)

func TestAppend(t *testing.T) {
	file := File{Buffer: "Hello World!\n", Adress: Range{2, 6}}
	newTest := "Big "
	file.Append(newTest)

	if file.Buffer != "Hello Big World!\n" {
		t.Error("expected ", "Hello Big World!", " got ", file.Buffer)
	}
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
	file := File{Buffer: "Hello World!\n", Adress: Range{6, 11}}
	file.Change("Moon")

	if file.Buffer != "Hello Moon!\n" {
		t.Errorf("expected '%s' got '%s'", "Hello Moon!", file.Buffer)
	}
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
	file := File{Buffer: "Hello World!\n", Adress: Range{5, 11}}
	file.DeleteDot()

	if file.Buffer != "Hello!\n" {
		t.Errorf("expected '%s' got '%s'", "Hello!", file.Buffer)
	}
	if file.Adress.Start != 5 || file.Adress.End != 5 {
		t.Error("range should have been set")
	}
}

func TestCopyDot(t *testing.T) {
	file := File{Buffer: "Hello World!", Adress: Range{5, 11}}
	file.CopyDot(5)

	if file.Buffer != "Hello World World!" {
		t.Errorf("expected '%s' got '%s'", "Hello World World!", file.Buffer)
	}
	if file.Adress.Start != 6 || file.Adress.End != 10 {
		t.Error("range should have been set")
	}
}

func BenchmarkCopyDot(b *testing.B) {
	file := File{Buffer: "Hello World!\n", Adress: Range{6, 11}}
	for i := 0; i < b.N; i++ {
		file.Append("Moon")
	}
}

func TestMoveDot(t *testing.T) {
	file := File{Buffer: "Hello World!\n", Adress: Range{6, 11}}
	file.Change("Moon")

	if file.Buffer != "Hello Moon!\n" {
		t.Errorf("expected '%s' got '%s'", "Hello Moon!", file.Buffer)
	}
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
