package main

import "fmt"

func main() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0

	//First way to loop through an array
	for i := 0; i < 5; i++ {
		total += x[i]
	}
	fmt.Println(total / 5)

	total = 0

	//Second way to loop through an array
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x))) //Type conversion

	total = 0

	//Third way to loop through an array
	//For each value in x, add it to total
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x))) //Type conversion

	x = [5]float64{98, 93, 77, 82, 83} //Array literal
	x = [5]float64{
		98,
		93,
		77,
		82,
		83, //Comma is required
	}
	fmt.Println(x)
}
