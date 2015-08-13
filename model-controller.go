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
	makes, err := mc.modelRepo.GetAll(MAX_MYSQL_RESULTS, nil)

	mc.Render(rw, makes, err)
}

func (mc ModelController) View(rw http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["id"])

	make, err := mc.modelRepo.Get(id, nil)

	mc.Render(rw, make, err)
}
