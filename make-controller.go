package main

import "net/http"

type MakeController struct {
	BaseController
	makeRepo IMakeRepository
}

func (mc MakeController) Index(rw http.ResponseWriter, request *http.Request) {

	resp := mc.makeRepo.GetAll(5)

	err := mc.Render(rw, resp)
	if err != nil {
		// TODO Implement as error method on base controller
		rw.Write([]byte("Internal server error"))
	}
}
