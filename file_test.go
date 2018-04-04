package main

import (
	"testing"
)

func TestAppend(t *testing.T) {
	file := File{Buffer: "Hello World!\n", Range: Range{2, 6}}
	file.Append("Big ")

	if file.Buffer != "Hello Big World!\n" {
		t.Error("Wrong appending result. Expected ", "Hello Big World!\n", " got ", file.Buffer)
	}
	if file.Range.Start != 6 || file.Range.End != 10 {
		t.Error("Range not properly updated after appending")
	}
}
