package main

import "fmt"

func main() {
	fmt.Println("Welcome to tutorial of array")

	//initializing the array
	// var array_name [size of array] array_data_type

	var branches [4]string
	branches[0] = "CSE"
	branches[1] = "ME"
	branches[3] = "MSE"

	fmt.Println("Branch list is :", branches)
	fmt.Println("Length of branch list is :", len(branches))                  //---->length of any variable or iterable can be measured with this len function
	fmt.Println("Length of code of mechanical branch is :", len(branches[1])) //---->length of any variable or iterable can be measured with this len function

	//initializing the array with default values
	var fruits = [3]string{"banana", "grapes", "mango"}

	fmt.Println("Fruit list is :", fruits)
	fmt.Println("Length of fruit list is :", len(fruits)) //---->length of any variable or iterable can be measured with this len function
}
