package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin) //------>bufio is the package used for std input output in GoLang
	fmt.Println("Please rate our pizza:")
	rate, err := reader.ReadString('\n')                               //----->kind of catch error statement till everything goes good value will be stored in rate and if something goes wrong error will be stores in err
	fmt.Println("Oh! wow thanks for rating, thank you so much", rate) //---->this line prints rating ( will always be printed in last )
	fmt.Printf("Type of this rating is %T \n", rate)
	fmt.Printf("%E", err) //----->error will be stored into err if somthing goes wrong and will be displayed through this line %T is for types and %E gives value
}
