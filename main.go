package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	PORT string = ":80"
)

func main() {
	// Instantiate our JsonRenderer for our api
	renderer := JsonRenderer{}

	// Instantiate BaseController with our JsonRenderer
	baseController := BaseController{
		IRenderer: &renderer,
	}

	// Instantiate repositories
	makeRepo := MakeMysqlRepo{}
	makeRepo.Init()

	defer func() {
		makeRepo.Deinit()
	}()

	// Instantiate application controllers
	makeController := MakeController{
		BaseController: baseController,
		makeRepo:       &makeRepo,
	}

	// Setup routes
	router := mux.NewRouter().StrictSlash(false)

	// GET /api/make
	router.HandleFunc("/api/make", makeController.Index).Methods("GET")

	log.Println("Listening on port 80...")
	http.ListenAndServe(PORT, router)
}
