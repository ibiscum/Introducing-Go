package main

import (
	"fmt"
	"time"
)

// The channel is represented by the chan keyword followed
// by the type of data that will be sent through the channel
func pinger(c chan string) {
	for i := 0; ; i++ {
		//The left arrow indicates that we are sending a message into the channel
		c <- "ping"
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c chan string) {
	for {
		//The right arrow indicates that we are receiving a message from the channel
		msg := <-c
		println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	//The channel is created with the make function
	c := make(chan string)

	//When pinger attempts to send a message on the channel
	//it will wait until printer is ready to receive the message.
	//This will call pinger and printer, ponger and printer.
	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
