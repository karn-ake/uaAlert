package controllers

import "net/http"

type Controllers interface {
	ClientController(res http.ResponseWriter, req *http.Request)
}
