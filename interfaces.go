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
		// Provide associated models as second param
		Get(int, []string) (Make, error)

		// GetAll returns all Make objects
		// Accepsts an int for limit
		// Provide associated models as second param
		GetAll(int, []string) ([]Make, error)

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
		// Provide associated models as second param
		Get(int, []string) (Model, error)

		// GetAll returns all Model objects
		// Accepsts an int for limit
		// Provide associated models as second param
		GetAll(int, []string) ([]Model, error)

		// Query returns a slice of Model objects
		// with custom string query and an error
		Query(string) ([]Model, error)
	}

	// ITagRepository interface
	// Enables hotswapping between database providers
	// if they implement the interface
	ITagRepository interface {
		// Get returns a single Tag struct object
		// with specified id or err on no results
		// Provide associated models as second param
		Get(int, []string) (Tag, error)

		// GetAll returns all Tag objects
		// Accepsts an int for limit
		// Provide associated models as second param
		GetAll(int, []string) ([]Tag, error)

		// Query returns a slice of Tag objects
		// with custom string query and an error
		Query(string) ([]Tag, error)
	}
)
