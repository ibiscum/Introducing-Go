package main

import "fmt"

//Variables

var globalValue = "Hello, World"

func main() {
	//Definition
	var x string
	x = "Hello, World"
	fmt.Println(x)
	x = "Second"
	fmt.Println(x)
	x = x + " Third"
	fmt.Println(x)

	//Equals
	var w string = "hello"
	var y = "hello" //Type inference
	fmt.Println(w == y)

	//Short definition
	a := "hello"
	fmt.Println(a)
	b := 5
	fmt.Println(b)

	dogsName := "Max" //Camel case
	fmt.Println("My dog's name is", dogsName)

	//Global variable
	fmt.Println(globalValue)

	const c string = "Hello, World" //Constants
	fmt.Println(c)

	//Multiple definitions
	var (
		l = 0
		u = 1
		i = 2
		z = 3
	)
	fmt.Println(l, u, i, z)
}

func f() {
	//Global variable
	fmt.Println(globalValue)
}
