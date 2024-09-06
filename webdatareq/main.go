package main

import (
	"fmt"
	"io"
	"net/url"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to webrequests tutorial")
	// getreq("https://dyord-backend.onrender.com/designforyou/events/?location=delhi")
	// postreq("https://dyord-backend.onrender.com/auth/login/")
	postformdata("https://dyord-backend.onrender.com/auth/login/")
}

func getreq(url string) {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Printf("Status Code:%v\n", response.StatusCode)
	fmt.Printf("Length:%v\n", response.ContentLength)

	content, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	var resString strings.Builder

	bytecount, err := resString.Write(content)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Bytecount is: %v\n", bytecount)
	fmt.Println(resString.String())
	// fmt.Println(content)
	// fmt.Println(string(content))
}
func postreq(url string) {
	reqbody := strings.NewReader(`
	{
		"email":"sanidhya738@gmail.com",
		"password":"009@NITHmsecoder"
	}
	`)

	response, err := http.Post(url, "application/json", reqbody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	var resbody strings.Builder
	content, _ := io.ReadAll(response.Body)
	bytecount, err := resbody.Write(content)
	if err != nil {
		panic(err)
	}
	fmt.Println(bytecount)
	fmt.Println(resbody.String())
}

func postformdata(url string) {
	dt := url.Values{}
	dt.Add("email", "sanidhya738@gmail.com")
	dt.Add("password", "009@NITHmsecoder")

	response, _ := http.PostForm(url, dt)

	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
