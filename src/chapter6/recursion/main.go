package main

func main() {
	// Calling a recursive function
	println(factorial(5))
}

// Recursive function
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
