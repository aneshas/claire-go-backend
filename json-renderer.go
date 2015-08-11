package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type (
	RespHead struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	Response struct {
		Head RespHead    `json:"head"`
		Body interface{} `json:"body,omitempty"`
	}

	JsonRenderer struct{}
)

// Implement IRenderer interface
func (jr JsonRenderer) Render(rw http.ResponseWriter, payload interface{}, queryErr error) (err error) {
	respHead := RespHead{}

	if queryErr != nil {
		respHead.Status = http.StatusNotFound
		respHead.Message = queryErr.Error()
		payload = nil
	} else {
		respHead.Status = http.StatusOK
		respHead.Message = "Ok"
	}

	response := Response{
		Head: respHead,
		Body: payload,
	}

	jsonPayload, err := json.Marshal(response)

	if err != nil {
		log.Println(err)
		return
	}

	rw.Header().Add("Content-Type", "application/json")
	_, err = rw.Write(jsonPayload)

	return
}
