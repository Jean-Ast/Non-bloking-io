package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// this function handles all http requests
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	//GET requests
	myRouter.Handle("/", http.HandlerFunc(homePage)).Methods("GET")
	myRouter.Handle("/", http.HandlerFunc(returnImage)).Methods("GET")

	//POST requests
	// myRouter.Handle("/", http.HandlerFunc()).Methods("POST")

	//PUT requests
	// myRouter.Handle("/}", http.HandlerFunc()).Methods("PUT")

	//DELETE requests
	// myRouter.Handle("/", http.HandlerFunc()).Methods("DELETE")

	// start server on port 10000 and binds to myRouter
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
func main() {
	fmt.Println("RESTful API up!")
	handleRequests()
}
