package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TagController struct {
	BaseController
	tagRepo ITagRepository
}

func (mc TagController) Index(rw http.ResponseWriter, request *http.Request) {
	tags, err := mc.tagRepo.GetAll(MAX_MYSQL_RESULTS, nil)

	mc.Render(rw, tags, err)
}

func (mc TagController) View(rw http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["id"])

	tag, err := mc.tagRepo.Get(id, nil)

	mc.Render(rw, tag, err)
}
