package main

import (
	"fmt"
	"introducingGo/src/chapter8/math"
)

//Alias
//import m "introducingGo/src/chapter8/math"

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
}
