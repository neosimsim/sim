package main

import "fmt"

// using string might be inefficient but should suffice for poc
var buffer string

type Range struct {
	Start int
	End int
}

var dotRange Range = Range {3,4}

func rangeContent(buffer string, r Range) string {
	return buffer[r.Start:r.End]
}

func main() {
	buffer = "Hallo Welt"
	fmt.Println("append:")
	fmt.Print(buffer[2:4])
	fmt.Println("change:")
	fmt.Print(buffer)
	fmt.Print(rangeContent(buffer, dotRange))
}

