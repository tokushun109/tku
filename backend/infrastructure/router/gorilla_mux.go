package router

import (
	"net/http"
	"os"

	"github.com/tokushun109/tku/backend/adapter/api/action"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type GorillaMuxServer struct {
	router  *mux.Router
	handler http.Handler
	port    Port
}

func NewGorillaMuxServer(port Port) *GorillaMuxServer {
	r := mux.NewRouter().StrictSlash(true)

	healthCheck := action.NewHealthCheckAction()
	r.HandleFunc("/api/health_check", healthCheck.Execute).Methods(http.MethodGet)

	clientURL := os.Getenv("CLIENT_URL")
	if clientURL == "" {
		panic("CLIENT_URL is required for CORS configuration")
	}

	c := cors.New(cors.Options{
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowedOrigins:   []string{clientURL},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	return &GorillaMuxServer{
		router:  r,
		handler: c.Handler(r),
		port:    port,
	}
}

func (g *GorillaMuxServer) Listen() {
	http.ListenAndServe(":"+string(g.port), g.handler)
}
