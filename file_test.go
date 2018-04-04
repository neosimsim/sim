package main

import (
	"testing"
)

func TestAppend(t *testing.T) {
	file := File{Buffer: "Hello World!\n", Dot: Range{2, 6}}
	file.Append("Big ")

	if file.Buffer != "Hello Big World!\n" {
		t.Error("Wrong appending result. Expected ", "Hello Big World!\n", " got ", file.Buffer)
	}
	if file.Dot.Start != 6 || file.Dot.End != 10 {
		t.Error("Dot not properly updated after appending")
	}
}
