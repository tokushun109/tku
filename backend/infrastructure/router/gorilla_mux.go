package router

import (
	"net/http"
	"os"

	"github.com/tokushun109/tku/backend/adapter/api/action"
	"github.com/tokushun109/tku/backend/adapter/api/middleware"
	"github.com/tokushun109/tku/backend/adapter/logger"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"gorm.io/gorm"
)

type GorillaMuxServer struct {
	router  *mux.Router
	handler http.Handler
	port    Port
	db      *gorm.DB
	log     logger.Logger
}

func NewGorillaMuxServer(log logger.Logger, db *gorm.DB, port Port) *GorillaMuxServer {
	r := mux.NewRouter().StrictSlash(true)

	healthCheck := action.NewHealthCheckAction(db, log)
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

	var handler http.Handler = r
	handler = middleware.NewRecovery(log)(handler)
	handler = middleware.NewLogger(log)(handler)
	handler = c.Handler(handler)

	return &GorillaMuxServer{
		router:  r,
		handler: handler,
		port:    port,
		db:      db,
		log:     log,
	}
}

func (g *GorillaMuxServer) Listen() {
	http.ListenAndServe(":"+string(g.port), g.handler)
}
