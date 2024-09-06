package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time function tutorial")
	presentTime := time.Now()
	fmt.Println("Unformated current time is \n", presentTime)
	fmt.Println("Formated current time is \n", presentTime.Format("02-Jan-2006 15:4/05PM Monday-->002"))
	createdDate := time.Date(2008,time.November,26,22,15,0,0,time.UTC)
	fmt.Println("Unformated created time is \n", createdDate)
	fmt.Println("Formated created time is \n", createdDate.Format("02-Jan-2006 3:4/05PM Monday-->002"))
}
// Year: "2006" "06"----->this is the format keys you have to use while structuring the format in which you need date of now
// Month: "Jan" "January" "01" "1"
// Day of the week: "Mon" "Monday"
// Day of the month: "2" "_2" "02"
// Day of the year: "__2" "002"
// Hour: "15" "3" "03" (PM or AM)
// Minute: "4" "04"
// Second: "5" "05"
// AM/PM mark: "PM"