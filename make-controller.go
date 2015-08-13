package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type MakeController struct {
	BaseController
	makeRepo IMakeRepository
}

func (mc MakeController) Index(rw http.ResponseWriter, request *http.Request) {
	makes, err := mc.makeRepo.GetAll(MAX_MYSQL_RESULTS, nil)

	err = mc.Render(rw, makes, err)
}

func (mc MakeController) View(rw http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["id"])

	queryParams := request.URL.Query()
	join := queryParams["join"]

	make, err := mc.makeRepo.Get(id, join)

	mc.Render(rw, make, err)
}
