package main

import "os"

func main() {

	// Open current directory
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()

	// Read directory
	//-1 means read all files and folders
	readDir, err := dir.ReadDir(-1)
	if err != nil {
		return
	}

	for _, file := range readDir {
		if file.IsDir() {
			println("Directory:", file.Name())
		} else {
			println("File:", file.Name())
		}
	}
}
