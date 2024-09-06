package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Welcome to Web request handling tutorial")

	url := "https://dyord-backend.onrender.com/designforyou/events/?location=delhi"

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Type of response is %T\n", res)
	defer res.Body.Close()
	databytes, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(databytes))

}
