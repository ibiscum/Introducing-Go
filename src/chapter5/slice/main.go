package main

import "fmt"

func main() {
	var x []float64

	//Making a slice with length 5
	x = make([]float64, 5)

	//A slice of length 5 with capacity of 10
	x = make([]float64, 5, 10)

	//Making a slice form an array [low:high]
	arr := [5]float64{1, 2, 3, 4, 5}
	x = arr[0:5]
	x = arr[0:] //is the same as arr[0:len(arr)]
	x = arr[:5] //is the same as arr[0:5]
	x = arr[:]  //is the same as arr[0:len(arr)]

	fmt.Println(x)

	//Append
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	slice2 = append(slice2, 6, 7)
	fmt.Println(slice1, slice2)

	//Copy
	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2) //Make a slice with length 2
	copy(slice4, slice3)     //Copy the first 2 elements of slice1 to slice2
	fmt.Println(slice3, slice4)
}
