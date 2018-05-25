package main

import (
	"testing"
)

func TestAppend(t *testing.T) {
	file := File{Buffer: "Hello World!\n", Range: Range{2, 6}}
	file.Append("Big ")

	if file.Buffer != "Hello Big World!\n" {
		t.Error("Wrong appending result. Expected ", "Hello Big World!", " got ", file.Buffer)
	}
	if file.Range.Start != 6 || file.Range.End != 10 {
		t.Error("Range not properly updated after appending.")
	}
}

func BenchmarkAppend(b *testing.B) {
	file := File{Buffer: "Hello World!\n", Range: Range{2, 6}}
	for i := 0; i < b.N; i++ {
		file.Append("Big ")
	}
}

func TestChange(t *testing.T) {
	file := File{Buffer: "Hello World!\n", Range: Range{6, 11}}
	file.Change("Moon")

	if file.Buffer != "Hello Moon!\n" {
		t.Error("Wrong change result. Expected ", "Hello Moon!", " got ", file.Buffer)
	}
	if file.Range.Start != 6 || file.Range.End != 10 {
		t.Error("Range not properly updated after changing dot.")
	}
}

func BenchmarkChange(b *testing.B) {
	file := File{Buffer: "Hello World!\n", Range: Range{6, 11}}
	for i := 0; i < b.N; i++ {
		file.Append("Moon")
	}
}

func TestDeleteDot(t *testing.T) {
	file := File{Buffer: "Hello World!\n", Range: Range{5, 11}}
	file.DeleteDot()

	if file.Buffer != "Hello!\n" {
		t.Error("Wrong deletion result. Expected ", "Hello!", " got ", file.Buffer)
	}
	if file.Range.Start != 5 || file.Range.End != 5 {
		t.Error("Range not properly updated after deleting dot.")
	}
}

func TestCopyDot(t *testing.T) {
	file := File{Buffer: "Hello World!\n", Range: Range{5, 11}}
	file.CopyDot(5)

	if file.Buffer != "Hello World World!\n" {
		t.Error("Wrong copy result. Expected ", "Hello Moon!", " got ", file.Buffer)
	}
	if file.Range.Start != 6 || file.Range.End != 10 {
		t.Error("Range not properly updated after changing dot.")
	}
}

func BenchmarkCopyDot(b *testing.B) {
	file := File{Buffer: "Hello World!\n", Range: Range{6, 11}}
	for i := 0; i < b.N; i++ {
		file.Append("Moon")
	}
}

func TestMoveDot(t *testing.T) {
	file := File{Buffer: "Hello World!\n", Range: Range{6, 11}}
	file.Change("Moon")

	if file.Buffer != "Hello Moon!\n" {
		t.Error("Wrong change result. Expected ", "Hello Moon!", " got ", file.Buffer)
	}
	if file.Range.Start != 6 || file.Range.End != 10 {
		t.Error("Range not properly updated after changing dot.")
	}
}

func BenchmarkMoveDot(b *testing.B) {
	file := File{Buffer: "Hello World!\n", Range: Range{6, 11}}
	for i := 0; i < b.N; i++ {
		file.Append("Moon")
	}
}

func TestDot(t *testing.T) {
	file := File{Buffer: "Hello World!\n", Range: Range{3, 7}}
	if file.Dot() != "lo W" {
		t.Error("Unexcpeted content of Dot. Expected ", "lo W", " got ", file.Dot())
	}
}
