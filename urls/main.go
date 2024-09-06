package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Welcome to tutorial on urls")

	myurl := "https://dyord-backend.onrender.com/designforyou/events/?location=delhi"

	parsed_url, _ := url.Parse(myurl)

	fmt.Println(parsed_url.Scheme)
	fmt.Println(parsed_url.Host)
	fmt.Println(parsed_url.Port())
	fmt.Println(parsed_url.Path)
	fmt.Println(parsed_url.RawQuery)

	qparams := parsed_url.Query()

	fmt.Println(qparams["location"])
	partsofurl := &url.URL{
		Scheme:   "https",
		Host:     "dyord-backend.onrender.com",
		Path:     "/designforyou/news/",
		RawQuery: "location=delhi",
	}
	madeurl := partsofurl.String()
	fmt.Println(madeurl)
}
