package routes

import (
	"log"
	"net/http"
	"uaAlert/controllers"

	"github.com/gorilla/mux"
)

type muxRoute struct {
	cont controllers.Controllers
}

var muxRouter = mux.NewRouter()

func New(cont controllers.Controllers) Routes {
	return &muxRoute{cont}
}

func (r *muxRoute) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxRouter.HandleFunc(uri, f).Methods("GET")
}

func (r *muxRoute) SERV(port string) {
	log.Println("MuxDispatcher's running on port:", port)
	http.ListenAndServe(port, muxRouter)
}
