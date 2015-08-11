package main

import "net/http"

type (
	// Response renderer interface
	// Should attempt to write[]byte to ResponseWritter,
	// returns nil on success, error on failure
	IRenderer interface {
		Render(http.ResponseWriter, interface{}, error) error
	}

	IDatabase interface {
		// Init initialize repository
		// Setup connections etc...
		Init()

		// Deinit - close db connectons etc...
		Deinit()
	}

	// IMakeRepository interface
	// Enables hotswapping between database providers
	// if they implement the interface
	IMakeRepository interface {
		// Get returns a single Make struct object
		// with specified id or err on no results
		Get(int) (Make, error)

		// GetAll returns all Make objects
		// Accepsts an int for limit
		GetAll(int) ([]Make, error)

		// Query returns a slice of Make objects
		// with custom string query and an error
		Query(string) ([]Make, error)
	}

	// IModelRepository interface
	// Enables hotswapping between database providers
	// if they implement the interface
	IModelRepository interface {
		// Get returns a single Model struct object
		// with specified id or err on no results
		Get(int) (Model, error)

		// GetAll returns all Model objects
		// Accepsts an int for limit
		GetAll(int) ([]Model, error)

		// Query returns a slice of Model objects
		// with custom string query and an error
		Query(string) ([]Model, error)
	}
)
