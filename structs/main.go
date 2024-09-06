package main

import "fmt"

func main() {
	fmt.Println("Welcome to tutorial on structs")
	//structs are basically an analogy to class
	//structs in golang do not have any inheritance
	sanidhya := User{"Sanidhya", "sanidhya738@gmail.com", 9736153960, true}
	fmt.Println(sanidhya)
	fmt.Printf("Type of object sanidhya is %T\n", sanidhya)
	fmt.Printf("Sanidhya details are here : %+v\n", sanidhya)
	fmt.Printf("My name is %v and my email is %v\n", sanidhya.Name, sanidhya.Email)

}

type User struct {
	Name    string
	Email   string
	Contact int
	Status  bool
}
