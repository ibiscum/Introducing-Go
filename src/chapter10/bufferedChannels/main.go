package main

import (
	"fmt"
	"time"
)

func main() {
	//Creation of a buffered channel with capacity 5
	ch := make(chan int, 5)

	// Pub Goroutine
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i // Send value to channel
			fmt.Println("Produzido:", i)
		}
		close(ch) // It is necessary to close the channel
	}()

	// Sub Goroutine
	go func() {
		for value := range ch {
			fmt.Println("Consumido:", value)
			time.Sleep(1 * time.Second) // Simulates a slow consumer
		}
	}()

	// Wait for the end of the program
	time.Sleep(15 * time.Second)
	fmt.Println("Fim do programa")
}
