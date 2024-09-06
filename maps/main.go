package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to tutorial about maps")

	languages := make(map[string]string)

	languages["rb"] = "Ruby"
	languages["js"] = "Javascript"
	languages["py"] = "Python"
	languages["cpp"] = "C++"

	fmt.Println(languages)
	fmt.Println(languages["rb"])
	fmt.Println(languages["py"])
	delete(languages, "cpp")
	fmt.Println(languages)
	fmt.Println("{")
	for key, value := range languages {
		fmt.Printf("languages[%v]:%v,\n", key, value)
	}
	fmt.Println("}")

}
