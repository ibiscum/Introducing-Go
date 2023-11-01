package main

import "os"

func main() {
	//ReadFile reads the file named by filename and returns the contents.
	bs, err := os.ReadFile("C:\\Projetos\\GoLang\\OReilly\\introducingGo\\src\\chapter8\\packages\\filesAndFolders\\test.txt")
	if err != nil {
		return
	}
	// Convert the byte slice to a string
	str := string(bs)
	println(str)
}
