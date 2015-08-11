package main

import "net/http"

type (
	// Response renderer interface
	// Should attempt to write[]byte to ResponseWritter,
	// returns nil on success, error on failure
	IRenderer interface {
		Render(http.ResponseWriter, interface{}) error
	}

	// IMakeRepository interface
	// Enables hotswapping between database providers
	// if they implement the interface
	IMakeRepository interface {
		// Init initialize repository
		// Setup connections etc...
		Init()

		// Deinit - close db connectons etc...
		Deinit()

		// Get returns a single Make struct object
		// with specified id, or nil
		Get(int) Make

		// GetAll returns all Make objects
		// Accepsts an int for limit (defaults to 20)
		GetAll(int) []Make

		// Query returns a slice of Make objects
		// with custom string query and an error
		Query(string) ([]Make, error)
	}
)
