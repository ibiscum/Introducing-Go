package main

import (
	"fmt"
	"os"
)

func main() {
	//Open the file
	file, err := os.Open("C:\\Projetos\\GoLang\\OReilly\\introducingGo\\src\\chapter8\\packages\\filesAndFolders\\test.txt")
	if err != nil {
		panic(err)
	}

	//Close the file after the function completes
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	//Get the file information
	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	//Create a buffer to hold the file's contents
	bs := make([]byte, stat.Size())

	//Read the file into the buffer
	_, err = file.Read(bs)
	if err != nil {
		panic(err)
	}

	//Convert the buffer into a string
	str := string(bs)

	fmt.Print(str)
}
