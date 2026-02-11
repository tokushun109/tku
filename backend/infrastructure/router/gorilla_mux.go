package router

import (
	"net/http"

	"github.com/tokushun109/tku/backend/adapter/api/action"

	"github.com/gorilla/mux"
)

type GorillaMuxServer struct {
	router *mux.Router
	port   Port
}

func NewGorillaMuxServer(port Port) *GorillaMuxServer {
	r := mux.NewRouter().StrictSlash(true)

	healthCheck := action.NewHealthCheckAction()
	r.HandleFunc("/api/health_check", healthCheck.Execute).Methods(http.MethodGet)

	return &GorillaMuxServer{router: r, port: port}
}

func (g *GorillaMuxServer) Listen() {
	http.ListenAndServe(":"+string(g.port), g.router)
}
