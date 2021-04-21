package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Image struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Src  string `json:"src"`
}

var Images []Image

func returnAllImages(w http.ResponseWriter, r *http.Request) {
	// Read image from file that already exists
	existingImageFile, err := os.Open("test.png")
	if err != nil {
		// Handle error
	}
	defer existingImageFile.Close()

	// Calling the generic image.Decode() will tell give us the data
	// and type of image it is as a string. We expect "png"
	imageData, _, err := image.Decode(existingImageFile)
	if err != nil {
		// Handle error
	}
	fmt.Println(imageData)
	// fmt.Println(imageType)

	// We only need this because we already read from the file
	// We have to reset the file pointer back to beginning
	existingImageFile.Seek(0, 0)

	// Alternatively, since we know it is a png already
	// we can call png.Decode() directly
	loadedImage, err := png.Decode(existingImageFile)
	if err != nil {
		// Handle error
	}
	fmt.Println(loadedImage)

	//Returns images
	// vars := mux.Vars(r)
	// key := vars["id"]

	// for _, image := range Images {
	// 	if image.Id == key {
	// 		json.NewEncoder(w).Encode(image)
	// 	}
	// }
	// json.NewEncoder(w).Encode(Images)
	fmt.Println("Endpoint Hit: returnImages")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// this function handles all http requests
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	//GET requests
	myRouter.Handle("/", http.HandlerFunc(homePage)).Methods("GET")
	myRouter.Handle("/api/images", http.HandlerFunc(returnAllImages)).Methods("GET")

	// start server on port 8080 and binds to myRouter
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
func main() {
	fmt.Println("RESTful API up!")
	Images = []Image{
		{Id: "1", Name: "Facebook", Src: "images\test.png"},
		{Id: "2", Name: "Mario", Src: "images\test2.png"},
	}
	handleRequests()
}
