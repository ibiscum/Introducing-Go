package main

import "fmt"

func main() {
	//Calling a variadic function
	fmt.Println(add(1, 2, 3))

	// Multiple arguments using a slice
	xs := []int{1, 2, 3}
	fmt.Println(add(xs...))
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}
