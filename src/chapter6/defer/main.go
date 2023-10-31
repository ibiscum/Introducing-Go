package main

import "fmt"

// Use defer to delay the execution of a statement until the end of the function
// It keeps our Close call near our Open call, so it's easier to understand
// If our function had multiple return statements (perhaps one in an if and one in an else) close will happen before both of them
// Deferred functions are run even if a runtime panic occurs
func main() {
	defer second()
	first()
}

func first() {
	fmt.Println("first")
}

func second() {
	fmt.Println("second")
}
