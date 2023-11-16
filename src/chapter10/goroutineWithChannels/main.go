package main

import (
	"fmt"
	"io"
	"net/http"
)

type HomePageSize struct {
	URL  string
	Size int
}

func main() {
	urls := []string{
		"https://www.apple.com/",
		"https://www.amazon.com/",
		"https://www.google.com/",
		"https://www.microsoft.com/",
	}

	results := make(chan HomePageSize)

	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()

			bs, err := io.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			results <- HomePageSize{
				URL:  url,
				Size: len(bs),
			}
		}(url)
	}

	var biggest HomePageSize

	//The for loop will wait for the results of the goroutines
	//Will execute the loop as many times as the number of urls
	for range urls {
		//Wait for the result of the goroutine
		result := <-results
		if result.Size > biggest.Size {
			biggest = result
		}
	}

	fmt.Println("The biggest home page:", biggest.URL)
}
