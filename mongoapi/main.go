package main

import (
	"fmt"
	"log"
	"net/http"

	router "github.com/sanidhya99/mongoapi/routers"
)

func main() {
	r := router.Router()
	fmt.Println("MongoDB API")
	fmt.Println("Server is getting started.....")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Server listening at port 4000.....")

}
