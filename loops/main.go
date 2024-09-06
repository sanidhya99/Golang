package main

import (
	"fmt"
	// "slices"
)

func main() {
	fmt.Println("Welcome to tutorial on loops")

	days := []string{"sunday", "monday", "tuesday", "wednusday", "thursday", "saturday"}

	// for i:=0;i<len(days);i=i+2{
	// 	fmt.Printf("Name of day %v is %v\n",i+1,days[i])
	// }

	// for i := range days{
	//   fmt.Println(days[i])
	// }

	// for index,value:=range days{
	// 	fmt.Printf("day corresponding to day%v is %v\n",index,value)
	// }
	for _, value := range days {
		fmt.Printf("day corresponding to day is %v\n", value)
	}
	limval := 1

	for limval < 10 {
		if limval == 2 {
			goto country
		}
		if limval == 3 {
			goto lco
		}
		if limval == 4 {
			goto trhee
		}

		if limval == 5 {
			break
		}
		fmt.Printf("Value is : %v\n", limval)
		limval++
	}

lco:
	fmt.Println("lco triggered")
country: //the label at which system gets through goto statement after that all code is executed including that inside triggers
	fmt.Println("JAI HIND !")
trhee:
	fmt.Println("3rd label trigerred")
	fmt.Println("2nd label was triggered successfully")
}
