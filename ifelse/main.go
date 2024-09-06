package main

import "fmt"

func main() {
	fmt.Println("Welcome to if-else statement tutorial")

	logincount := 100
	var result string
	//syntax2
	if logincount < 10 {
		result = "regular user"
	} else if logincount > 10 {
		result = "binge watcher"
	} else {
		result = "medium watch"
	}
	//syntax2
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Even number it is")
	} else {
		fmt.Println("Odd number it is")
	}
	//syntax3
	if num := 3; num < 10 {
		fmt.Println("num is smaller")
	} else if num > 10 {
		fmt.Println("num is greater")
	} else {
		fmt.Println("num is equal")
	}
}
