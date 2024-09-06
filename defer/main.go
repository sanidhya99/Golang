package main

import "fmt"

func main() {
	defer fmt.Println("People")
	defer fmt.Println("Golang")
	defer fmt.Println("World")
	// fmt.Println("World Golang people")
	fmt.Println("Hello")
	deferPr()
}

func deferPr() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

//when code is executed in golang it keep putting all defer statements in a stack and keeps executing all normal statements and later just at the end of function when function is just about to return at that time all defer statements are executed in Li-Fo form from stack.
