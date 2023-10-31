package main

import "fmt"

func main() {
	// Closure, define a function inside another function
	add := func(x, y int) int {
		return x + y
	}
	println(add(1, 1))

	x := 0
	increment := func() int {
		x++ // x is defined in the outer function
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())

	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) //0
	fmt.Println(nextEven()) //2
	fmt.Println(nextEven()) //4
}

// This is a function that returns another function
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}
