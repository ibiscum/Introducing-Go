package main

import "fmt"

func main() {
	x := 5
	zeroWithoutPointer(x)
	fmt.Println(x) // x is still 5

	//We use the & operator to find the address of a variable. &x returns a *int (pointer to an int) because x is an int.
	zeroWithPointer(&x)
	fmt.Println(x) // x is 0 now

	//The "new" takes a type as an argument, allocates enough memory to fit a value of that type and returns a pointer to it.
	y := new(int)
	one(y)
	fmt.Println(*y) // y is 1 now
}

func zeroWithoutPointer(x int) {
	//In this case we are passing a copy of the value of x, so the function will not change the value of x
	x = 0
}

func zeroWithPointer(x *int) {
	//In Go a pointer is represented using the * (asterisk) character followed by the type of the stored value.
	*x = 0 //If whe remove the *, we wil get a compile error, because we are trying to assign a value to a pointer
}

func one(x *int) {
	*x = 1
}
