package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/gorilla/mux"
	//sw "./go/"
)

func main() {

	fmt.Println("Hello, world.")

	log.Printf("Server started")

	r := mux.NewRouter()

	//router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", r))

}
