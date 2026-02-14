package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tokushun109/tku/clean-backend/internal/infra/config"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/handler"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/middleware"
)

func NewRouter(cfg *config.Config, categoryHandler *handler.CategoryHandler, auth *middleware.AuthMiddleware) http.Handler {
	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)

	r.HandleFunc("/api/category", categoryHandler.List).Methods(http.MethodGet)
	r.Handle("/api/category", auth.RequireSession(http.HandlerFunc(categoryHandler.Create))).Methods(http.MethodPost)
	r.Handle("/api/category/{category_uuid}", auth.RequireSession(http.HandlerFunc(categoryHandler.Update))).Methods(http.MethodPut)
	r.Handle("/api/category/{category_uuid}", auth.RequireSession(http.HandlerFunc(categoryHandler.Delete))).Methods(http.MethodDelete)

	return middleware.CORSMiddleware([]string{cfg.ClientURL})(r)
}
