package main

import (
	"fmt"
	"math"
)

func main() {
	//We can create a struct in two ways:

	//1. Using the "var" keyword
	//This will set zero values for all the fields
	var c1 Circle

	//2. Using the "new" keyword
	//This will alocate memory for all the fields and return a pointer to the struct
	c2 := new(Circle)

	//We can also initialize the fields of the struct
	c3 := Circle{x: 0, y: 0, r: 5}
	c4 := Circle{0, 0, 5}

	//We can also create a pointer to a struct using the & operator
	//This will create de pointer and initialize the fields
	c5 := &Circle{x: 0, y: 0, r: 5}

	//We can access the fields of a struct using the . operator
	c1.x = 10
	c1.y = 5
	c1.r = 2.5

	fmt.Println(c1, c2, c3, c4, c5)

	c := Circle{0, 0, 5}
	fmt.Println(circleArea(c))

	fmt.Println(c.area())
}

// Struct
type Circle struct {
	x, y, r float64
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

// Methods
// This is as an extension method in c#
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}
