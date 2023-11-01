package main

import (
	"fmt"
	"strings"
)

func main() {
	// Contains Check if a string contains a substring
	containsResult := strings.Contains("Hello World", "Hello")
	fmt.Println(containsResult) // true

	//Count the number of non-overlapping instances of a substring in a string
	countResult := strings.Count("Hello World", "l")
	fmt.Println(countResult) //3

	//HasPrefix Check if a string starts with a prefix
	hasPrefixResult := strings.HasPrefix("Hello World", "Hello")
	fmt.Println(hasPrefixResult) // true

	//HasSuffix Check if a string ends with a suffix
	hasSuffixResult := strings.HasSuffix("Hello World", "World")
	fmt.Println(hasSuffixResult) // true

	//Index find the index of the first instance of a substring in a string
	indexResult := strings.Index("Hello World", "World")
	fmt.Println(indexResult) // 6

	//Join a slice of strings into a single string using a separator
	joinResult := strings.Join([]string{"Hello", "World"}, " ")
	fmt.Println(joinResult) // Hello World

	//Repeat a string n times
	repeatResult := strings.Repeat("Hello", 3)
	fmt.Println(repeatResult) // HelloHelloHello

	//Replace all instances of one substring with another
	replaceResult := strings.Replace("Hello World", "World", "Go", 1)
	fmt.Println(replaceResult) // Hello Go

	//Split a string into a slice of substrings using a separator
	splitResult := strings.Split("a-b-c-d-e", "-")
	fmt.Println(splitResult) // [a b c d e]

	//ToLower Convert a string to lowercase
	toLowerResult := strings.ToLower("Hello World")
	fmt.Println(toLowerResult) // hello world

	//ToUpper Convert a string to uppercase
	toUpperResult := strings.ToUpper("Hello World")
	fmt.Println(toUpperResult) // HELLO WORLD

	//Trim leading and trailing whitespace from a string
	trimResult := strings.TrimSpace(" Hello World ")
	fmt.Println(trimResult) // Hello World

	//Trim leading and trailing characters specified in a string from a string
	trimResult = strings.Trim("...Hello World...", ".")
	fmt.Println(trimResult) // Hello World

	//Trim leading characters specified in a string from a string
	trimResult = strings.TrimLeft("...Hello World...", ".")
	fmt.Println(trimResult) // Hello World...

	//Trim trailing characters specified in a string from a string
	trimResult = strings.TrimRight("...Hello World...", ".")
	fmt.Println(trimResult) // ...Hello World

	//Convert a slice of bytes to a string
	bytes := []byte{'H', 'e', 'l', 'l', 'o'}
	stringResult := string(bytes)
	fmt.Println(stringResult) // Hello
}
