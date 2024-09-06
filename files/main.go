package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to the tutorial about handling files")

	file, err := os.Create("./veergathadoc.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilerr(err)

	content := "It is a website which brings unheard gems to people"

	length, err := io.WriteString(file, content)

	checkNilerr(err)
	defer file.Close()
	fmt.Println("length of file is ", length)
	readF("./veergathadoc.txt")
}

func readF(file string) {
	databyte, err := os.ReadFile(file) //it returns content of file in bytes format
	checkNilerr(err)
	fmt.Println(string(databyte))

}

func checkNilerr(err error) {
	if err != nil {
		panic(err)
	}
} //std. error handling in golang
