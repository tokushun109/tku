package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/handler"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/middleware"
)

func NewRouter(
	healthHandler *handler.HealthHandler,
	categoryHandler *handler.CategoryHandler,
	targetHandler *handler.TargetHandler,
	auth *middleware.AuthMiddleware,
	logging func(http.Handler) http.Handler,
	cors func(http.Handler) http.Handler,
) http.Handler {
	r := mux.NewRouter()
	r.Use(logging)

	r.HandleFunc("/api/health_check", healthHandler.Check).Methods(http.MethodGet)
	r.HandleFunc("/api/category", categoryHandler.List).Methods(http.MethodGet)
	r.HandleFunc("/api/target", targetHandler.List).Methods(http.MethodGet)
	// TODO: User取得処理作成後、管理者権限チェック（AdminMiddleware 等）を追加する。
	r.Handle("/api/category", auth.RequireSession(http.HandlerFunc(categoryHandler.Create))).Methods(http.MethodPost)
	r.Handle("/api/category/{category_uuid}", auth.RequireSession(http.HandlerFunc(categoryHandler.Update))).Methods(http.MethodPut)
	r.Handle("/api/category/{category_uuid}", auth.RequireSession(http.HandlerFunc(categoryHandler.Delete))).Methods(http.MethodDelete)
	r.Handle("/api/target", auth.RequireSession(http.HandlerFunc(targetHandler.Create))).Methods(http.MethodPost)
	r.Handle("/api/target/{target_uuid}", auth.RequireSession(http.HandlerFunc(targetHandler.Update))).Methods(http.MethodPut)
	r.Handle("/api/target/{target_uuid}", auth.RequireSession(http.HandlerFunc(targetHandler.Delete))).Methods(http.MethodDelete)

	return cors(r)
}
