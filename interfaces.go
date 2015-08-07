package main

import "net/http"

type (
	// Response renderer interface
	// Should attempt to write[]byte to ResponseWritter,
	// returns nil on success, error on failure
	IRenderer interface {
		Render(http.ResponseWriter, interface{}) error
	}
)
