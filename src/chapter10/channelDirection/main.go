package main

import (
	"fmt"
	"time"
)

//Now we'll look at channel directions. When using channels as function parameters,
//you can specify if a channel is meant to only send or receive values. This specificity
//increases the type-safety of the program.

// Setting chan<- as the parameter to ping enables us to pass only a channel that is meant to SEND data.
func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

// Setting <-chan as them parameter to printer enables us to pass only a channel that is meant to RECEIVE data.
func printer(c <-chan string) {
	for {
		msg := <-c
		println(msg)
		time.Sleep(time.Second * 1)
	}
}

// Channels that doesn't have these restrictions is known as bidirectional.

func main() {
	c := make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
