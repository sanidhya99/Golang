package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to tutorial about mod")
	greeter()
	r := mux.NewRouter()

	r.HandleFunc("/", homeserve).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}
func greeter() {
	fmt.Println("Hey mod users welcome to this tutorial")
}
func homeserve(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello Golang Backend Developers its our first API in golang"))
}
