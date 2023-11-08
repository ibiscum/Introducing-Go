package main

import (
	"errors"
	"fmt"
)

func main() {
	// Create an error
	err := errors.New("error message")

	fmt.Print(err)
}
