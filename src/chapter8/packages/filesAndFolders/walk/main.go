package main

import (
	"os"
	"path/filepath"
)

func main() {
	// Walk current directory
	// This will walk the current directory and all subdirectories
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			println("Directory:", path)
		} else {
			println("File:", path)
		}
		return nil
	})
}
