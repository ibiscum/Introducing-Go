package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)

		//The rand.Intn function returns a random int n, 0 <= n < 100
		amt := time.Duration(rand.Intn(250))

		//The time.Sleep function pauses the execution of the current goroutine for at least the duration d
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		//The keyword go represent a goroutine
		go f(i)
	}

	fmt.Println("Press enter to continue...")

	//The Scanln function from the fmt package will wait for the user to press enter before the program exits
	//This is a hack to keep the program running so that the goroutines have a chance to finish
	var input string
	fmt.Scanln(&input)
}
