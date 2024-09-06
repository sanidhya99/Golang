package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to tutorial of slices")
	//syntax to initialize slice

	var city = []string{} //you dont need to specify number of elements , you can add as many elements as you want memory will automatically expand

	fmt.Println(city)
	fmt.Println(len(city))
	fmt.Printf("Type of data is : %T \n", city)
	// city=append(city, "nahan")
	// fmt.Println(len(city))
	// city=append(city, "Delhi")
	// fmt.Println(len(city))
	// city=append(city, "Chandigarh")
	// fmt.Println(len(city))
	// city=append(city, "Dharamshala")
	// fmt.Println(len(city))

	//Now there are 2 methods to add elements in slice
	//1. you can add some elements at time of intializing
	var marks = []int{1, 2, 3, 5}
	fmt.Println(marks)
	//2. you can also add elements later on by append method
	// slice_name=append(slice_name,element1_to_be_added,element2_to_be_added)
	marks = append(marks, 10)
	fmt.Println(marks)

	//Slicing of data

	marks = append(marks[1:4]) //start index inclusive , end index exclusive , if you leave any value from 2 blank like marks[:4] or marks[1:] so it will take that as end indexes start index as 0 and end index as ( length of slice - 1 )
	fmt.Println(marks)

	me := "Sanidhya"
	fmt.Println(me[1:3]) //strings can also be sliced in golang in same way

	//A different way to initialize a data of any data type

	speeds := make([]int, 4) //now this will make a slice speeds of int data type initializing 4 elements with no value
	fmt.Println(speeds)
	fmt.Println(len(speeds))
	fmt.Printf("Type of data is : %T\n", speeds)
	speeds[0] = 20
	speeds[1] = 40
	speeds[2] = 650
	speeds[3] = 800
	fmt.Println(speeds)
	fmt.Println(len(speeds))
	fmt.Printf("Type of data is : %T\n", speeds)
	//slice or any data type made with make function is alloted memory only equal to the value mentioned in function that is 4 after that if you want to extend slice you can use append function

	speeds = append(speeds, 100, 150)

	fmt.Println(speeds)
	fmt.Println(len(speeds))
	fmt.Printf("Type of data is : %T\n", speeds)
	fmt.Println(sort.IntsAreSorted(speeds))
	sort.Ints(speeds) //through sorts function you can sort any kind of data type slices, ints , strings etc it aslo have bool functions which can tell whether the data of particular data type is sorted or not

	fmt.Println(speeds)
	fmt.Println(sort.IntsAreSorted(speeds))

	var courses = []string{"js", "python", "java", "c++", "golang", "django", "aws"}

	ind := 2
	courses = append(courses[:ind], courses[ind+1:]...) //in last ... is important as in append first argument in slice want to edit but after that it wants all arguments to be of data type same as that of slice but here we are giving slice instead of string so we have to give this ... to make it valid syntax
	fmt.Println(courses)

	// var refsl=[]string{}
	// refsl=append(courses[:ind],courses[ind+1:]...)
	// fmt.Println(refsl)
}
