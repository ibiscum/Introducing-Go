package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	//create a hasher
	//crc32.NewIEEE() creates a new hash.Hash32 object
	//which implements the hash.Hash interface
	//the IEEE polynomial is used in the CRC-32 algorithm
	//IEEE is the Institute of Electrical and Electronics Engineers
	//the IEEE is a professional association that develops
	//standards for the computer and electronics industry
	h := crc32.NewIEEE()

	//write our data to it
	//the crc32 hash object implements the writer interface
	_, err := h.Write([]byte("test"))
	if err != nil {
		return
	}

	//calculate the crc32 checksum
	//A common use for crc32 is to compare two files
	//by comparing their checksums
	v := h.Sum32()
	fmt.Println(v)
}
