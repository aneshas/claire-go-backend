package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ModelController struct {
	BaseController
	modelRepo IModelRepository
}

func (mc ModelController) Index(rw http.ResponseWriter, request *http.Request) {
	makes, err := mc.modelRepo.GetAll(2)

	err = mc.Render(rw, makes, err)
	if err != nil {
		// TODO Implement as error method on base controller
		rw.Write([]byte("Internal server error"))
	}
}

func (mc ModelController) View(rw http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["id"])

	make, err := mc.modelRepo.Get(id)

	mc.Render(rw, make, err)
}
