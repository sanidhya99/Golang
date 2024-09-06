package main

//In this file we will get to know about declaration and use of variables
import "fmt"

// username := "cptbrandon" - you can not use no var type outside func main()
var Username = "cptbrandon"

func main() {
	//declaring int variable
	var users int = 8
	fmt.Println(users)
	fmt.Printf("Type of variable is : %T \n", users)

	//declaring int variable
	var ref int
	fmt.Println(ref)//if you initialize int variable without specifying initial value , then it will be initialized with default value 0 
	fmt.Printf("Type of variable is : %T \n", ref)

	//declaring string variable
	var message string = "Golang is good language"
	fmt.Println(message)
	fmt.Printf("Type of variable is : %T \n", message)
	//declaring bool variable
	var isActive bool = true
	fmt.Println(isActive)
	fmt.Printf("Type of variable is : %T \n", isActive)
	//declaring float variable
	var salary float32 = 400000.423
	fmt.Println(salary)
	fmt.Printf("Type of variable is : %T \n", salary)

	//implicit type declaration
	var weburl = "www.veergatha.netlify.app/"
	fmt.Println(weburl)
	fmt.Printf("Type of variable is : %T \n", weburl)
	// weburl = 1 --> in implicit declaration you can not redeclare the variable with another datatype however you can update value of varible with same datatype

	//free style(no var style declaration)
	jwtToken := "token_65e#4534"
	fmt.Println(jwtToken)
	fmt.Printf("Type of variable is : %T \n", jwtToken)
	// jwtToken = 8 --> in no var style declaration you can not redeclare the variable with another datatype however you can update value of varible with same datatype

	fmt.Println(Username)
	fmt.Printf("Type of variable is : %T \n", Username)

	const slogan = "naam namak nishan"
	fmt.Println(slogan)
	fmt.Printf("Type of variable is : %T \n", slogan)

	// slogan = "Jai Hind" --> const entities can not be updated with new values , they are constant .
}
