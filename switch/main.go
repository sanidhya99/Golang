package main

import "fmt"

func main() {
	fmt.Println("Welcome to tuorial of switch case statements.")

	loggedIn := "yes"

	switch loggedIn {
	case "yes":
		fmt.Println("Authenticated")
	case "no":
		fmt.Println("Unathenticated")
	default:
		fmt.Println("Server Error 500")
	}

	turnOver := 300

	switch turnOver {
	case 100:
		fmt.Println("Starter")
	case 200:
		fmt.Println("Good going")
	case 300:
		fmt.Println("Unicorn")
	default:
		fmt.Println("404 Not Found")
	}

}
