package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			//This works like a swicth statement, but for channels
			//It will select the first channel that is ready and receive from it
			//If more than one channel is ready then it will randomly select which one to receive from
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			//time.after creates a channel and after the given duration will send the current time on it
			case <-time.After(time.Second):
				fmt.Println("timeout")
			//The default case happens immediately if none of the channels are ready
			default:
				fmt.Println("nothing ready")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
