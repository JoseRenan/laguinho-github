package app

import (
	"log"
	"net/http"

	"github.com/JoseRenan/laguinho-github/api"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	addr   string
}

func (a *App) NewApp(addr string) *App {
	a.Router = mux.NewRouter()
	a.addr = addr
	return a
}

func (a *App) initializeRoutes() {
	a.Router.
		HandleFunc("/datasets/{owner}/{repo}", api.GetDataset).
		Queries("path", "{path}").
		Methods(http.MethodGet)
}

func (a *App) Run() {
	log.Println("Listening to", a.addr)
	a.initializeRoutes()
	log.Fatal(http.ListenAndServe(a.addr, a.Router))
}
