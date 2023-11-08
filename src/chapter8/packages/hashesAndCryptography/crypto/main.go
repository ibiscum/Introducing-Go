package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	//Create a hasher
	h := sha1.New()

	//Write out data to it
	//the sha1 hash object implements the writer interface
	h.Write([]byte("test"))

	//Sha1 cimputes a 160-bit hash
	bs := h.Sum([]byte{})

	fmt.Println(bs)
}
