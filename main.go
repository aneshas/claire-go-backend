package main

import (
	"fmt"
	"net/http"
)

const (
	PORT string = ":80"
)

func main() {
	// Instantiate JsonRenderer for our api
	jsonRenderer := JsonRenderer{}

	// Instantiate BaseController with our JsonRenderer
	baseController := BaseController{
		IRenderer: &jsonRenderer,
	}

	// Instantiate repositories

	// Instantiate application controllers
	makeController := MakeController{
		BaseController: baseController,
	}

	// Setup routes
	http.HandleFunc("/", makeController.Index)

	fmt.Println("Listening on port 80...")
	http.ListenAndServe(PORT, nil)
}
