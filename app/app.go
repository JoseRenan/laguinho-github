package app

import (
	"github.com/JoseRenan/laguinho-github/api"
	"github.com/gorilla/mux"
	"net/http"
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
	a.initializeRoutes()
	http.ListenAndServe(a.addr, a.Router)
}
