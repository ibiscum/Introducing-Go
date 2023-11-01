package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	//Buffer does not need to be initialized
	var buf bytes.Buffer

	//Write to the buffer
	buf.Write([]byte("Hello "))

	x := "World"
	//Create a reader from a string
	reader := strings.NewReader(x)

	//Write the contents of the reader to the buffer
	_, err := reader.WriteTo(&buf)
	if err != nil {
		return
	}

	fmt.Println(buf.String()) // Hello World
}
