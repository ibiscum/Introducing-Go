package main

import "os"

func main() {
	// Create creates or truncates the named file.
	file, err := os.Create("C:\\Projetos\\GoLang\\OReilly\\introducingGo\\src\\chapter8\\packages\\filesAndFolders\\create\\test.txt")
	if err != nil {
		return
	}

	defer file.Close()

	// WriteString writes the contents of the string s to the file, returning the number of bytes written.
	_, err = file.WriteString("Hello World")
	if err != nil {
		return
	}
}
