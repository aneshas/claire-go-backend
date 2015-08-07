package main

import "net/http"

type MakeController struct {
	BaseController
}

func (c MakeController) Index(rw http.ResponseWriter, request *http.Request) {
	resp := []Make{
		{
			Id:          1,
			Name:        "BMW",
			Description: "BMW Description",
			Photo:       "A great BMW photo",
		},
		{
			Id:          1,
			Name:        "Audi",
			Description: "Audi Description",
			Photo:       "A great Audi photo",
		},
	}

	c.Render(rw, resp)
}
