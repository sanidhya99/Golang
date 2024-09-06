package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	urls := []string{
		"https://www.google.com",
		"https://veergatha.netlify.app",
		"https://www.goyuogle.com",
	}
	for _, web := range urls {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
}

func getStatusCode(url string) {
	defer wg.Done()
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("OOPS in url")
	} else {
		fmt.Printf("%d status code for %s\n", res.StatusCode, url)
	}
}
