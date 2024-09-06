package main

import "fmt"

func main() {
	fmt.Println("Welcome to tutorial of functions")
	sum, st, pr := app(4, 6)
	fmt.Printf("%v,%v,%v", sum, st, pr)
}

func app(aval int, bval int) (int, string, int) {
	return (aval + bval), ("OK Added"), (aval * bval)
}
