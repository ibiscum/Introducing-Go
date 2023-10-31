package main

import "fmt"

func main() {
	//Integer Types
	// uint - unsigned integer
	// int - signed integer
	// byte - alias for uint8
	// rune - alias for int32
	// uint, int, and uintptr depend on the platform
	fmt.Println("1 + 1 =", 1+1)

	//Floating Point Types
	//Are inexact
	// float32 - 32 bits
	// float64 - 64 bits
	// NaN and Infinity are possible values

	//Complex Types
	// complex64 - float32 real and imaginary parts
	// complex128 - float64 real and imaginary parts
	fmt.Println("1 + 1 =", 1.0+1.0)

	//Strings
	// UTF-8
	// Immutable
	// Can be concatenated with +
	// Can be converted to []byte
	// Can be indexed, with the result of the index being the byte, not the character
	fmt.Println("Hello, World")
	fmt.Println(`Hello, 
	World`) // backticks allow for multi-line strings
	fmt.Println(len("Hello, World"))
	fmt.Println("Hello, World"[1])
	fmt.Println("Hello, " + "World")

	//Booleans
	// true
	// false
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
