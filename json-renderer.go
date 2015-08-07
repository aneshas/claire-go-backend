package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type JsonRenderer struct{}

// Implement IRenderer interface
func (jr JsonRenderer) Render(rw http.ResponseWriter, payload interface{}) (err error) {
	jsonPayload, err := json.Marshal(payload)

	if err != nil {
		log.Println(err)
		return
	}

	rw.Header().Add("Content-Type", "application/json")
	_, err = rw.Write(jsonPayload)

	return
}
