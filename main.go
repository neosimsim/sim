package main

import "fmt"

// using string might be inefficient but should suffice for poc
var buffer string

var file File = File{
	Buffer: "Hallo Welt",
	Adress: Range{4, 8},
}

var dotRange Range = Range{3, 4}

func main() {
	buffer = "Hallo Welt"
	fmt.Println("append:")
	fmt.Print(buffer[2:4])
	fmt.Println("change:")
	fmt.Print(buffer)
	fmt.Print(file.Dot())
}
