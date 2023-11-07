package main

import "fmt"

//Variables

var globalValue = "Hello, World"

func main() {
	//Definition
	var x string
	x = "Hello, World"
	fmt.Println(x) // Hello, World
	x = "Second"
	fmt.Println(x) // Second
	x = x + " Third"
	fmt.Println(x) // Second Third

	//Equals
	var w string = "hello"
	var y = "hello"     //Type inference
	fmt.Println(w == y) // true

	//Short definition
	a := "hello"
	fmt.Println(a) // hello
	b := 5
	fmt.Println(b) // 5

	dogsName := "Max"                         //Camel case
	fmt.Println("My dog's name is", dogsName) // My dog's name is Max

	//Global variable
	fmt.Println(globalValue) // Hello, World

	const c string = "Hello, World" //Constants
	fmt.Println(c)                  // Hello, World

	//Multiple definitions
	var (
		l = 0
		u = 1
		i = 2
		z = 3
	)
	fmt.Println(l, u, i, z) // 0 1 2 3
}

func f() {
	//Global variable
	fmt.Println(globalValue) // Hello, World
}
