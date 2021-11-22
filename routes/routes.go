package routes

import "net/http"

type Routes interface {
	GET(uri string, f func(w http.ResponseWriter, r *http.Request))
	SERV(port string)
}
