package main

import "fmt"

func main() {

	//While
	i := 1
	for i <= 10 {
		fmt.Println(i) //Prints 1 to 10
		i += 1
	}

	//For
	for j := 1; j <= 10; j++ {
		fmt.Println(j)
	}

	//For each
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83
	for i, value := range x {
		fmt.Println(i, value)
	}
}
