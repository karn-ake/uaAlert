package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"uaAlert/repository"
	"uaAlert/services"

	"github.com/gorilla/mux"
)

type muxController struct {
	serv services.Services
	repo repository.Repository
}

func New(serv services.Services, repo repository.Repository) MuxControllers {
	return &muxController{serv, repo}
}

func (c *muxController) ClientController(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	cn := strings.ToUpper(mux.Vars(req)["client"])
	fn, _ := c.repo.FindbyClientName(cn)
	post, err := c.serv.CheckStatus(cn, fn.LogFile)
	if err != nil {
		log.Println(err)
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(&post)
}
