package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/handler"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/middleware"
)

// TODO: User取得処理作成後、管理者権限チェック（AdminMiddleware 等）を追加する
func NewRouter(
	healthHandler *handler.HealthHandler,
	categoryHandler *handler.CategoryHandler,
	targetHandler *handler.TargetHandler,
	tagHandler *handler.TagHandler,
	salesSiteHandler *handler.SalesSiteHandler,
	auth *middleware.AuthMiddleware,
	logging func(http.Handler) http.Handler,
	cors func(http.Handler) http.Handler,
) http.Handler {
	r := mux.NewRouter().StrictSlash(true)
	r.Use(logging)

	// health check
	r.HandleFunc("/api/health_check", healthHandler.Check).Methods(http.MethodGet)

	// category
	r.HandleFunc("/api/category", categoryHandler.List).Methods(http.MethodGet)
	r.Handle("/api/category", auth.RequireSession(http.HandlerFunc(categoryHandler.Create))).Methods(http.MethodPost)
	r.Handle("/api/category/{category_uuid}", auth.RequireSession(http.HandlerFunc(categoryHandler.Update))).Methods(http.MethodPut)
	r.Handle("/api/category/{category_uuid}", auth.RequireSession(http.HandlerFunc(categoryHandler.Delete))).Methods(http.MethodDelete)

	// target
	r.HandleFunc("/api/target", targetHandler.List).Methods(http.MethodGet)
	r.Handle("/api/target", auth.RequireSession(http.HandlerFunc(targetHandler.Create))).Methods(http.MethodPost)
	r.Handle("/api/target/{target_uuid}", auth.RequireSession(http.HandlerFunc(targetHandler.Update))).Methods(http.MethodPut)
	r.Handle("/api/target/{target_uuid}", auth.RequireSession(http.HandlerFunc(targetHandler.Delete))).Methods(http.MethodDelete)

	// tag
	r.HandleFunc("/api/tag", tagHandler.List).Methods(http.MethodGet)
	r.Handle("/api/tag", auth.RequireSession(http.HandlerFunc(tagHandler.Create))).Methods(http.MethodPost)
	r.Handle("/api/tag/{tag_uuid}", auth.RequireSession(http.HandlerFunc(tagHandler.Update))).Methods(http.MethodPut)
	r.Handle("/api/tag/{tag_uuid}", auth.RequireSession(http.HandlerFunc(tagHandler.Delete))).Methods(http.MethodDelete)

	// sales site
	r.HandleFunc("/api/sales_site", salesSiteHandler.List).Methods(http.MethodGet)
	r.Handle("/api/sales_site", auth.RequireSession(http.HandlerFunc(salesSiteHandler.Create))).Methods(http.MethodPost)
	r.Handle("/api/sales_site/{sales_site_uuid}", auth.RequireSession(http.HandlerFunc(salesSiteHandler.Update))).Methods(http.MethodPut)
	r.Handle("/api/sales_site/{sales_site_uuid}", auth.RequireSession(http.HandlerFunc(salesSiteHandler.Delete))).Methods(http.MethodDelete)

	return cors(r)
}
