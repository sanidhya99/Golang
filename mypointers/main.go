package main

import "fmt"

func main() {
	fmt.Println("Welcome to tutorial on mypointers")
	var one uint16 = 1

	// var two = 2
	// three := 3
	//pointers in go also haave exactly same concept as in other languages

	fmt.Println("numbers", &one)
	fmt.Println("numbers", *&one)
	var ptr = &one
	*ptr = *ptr * 10
	fmt.Println("numbers", one)
}
