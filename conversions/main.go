package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Hello welcome to our app! \n Please rate our services b/w 1 to 5")

	rate, _ := reader.ReadString('\n')

	fmt.Println("Oh ! Thanks for such a good rating, ", rate)

	fmt.Printf("The type of this rating is %T \n", rate)

	numrate, _ := strconv.ParseInt(strings.TrimSpace(rate), 0, 64)

	fmt.Println("We have updated your rating to : ", numrate+1)

	fmt.Printf("%T", numrate)

}
