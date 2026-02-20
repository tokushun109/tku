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
	tagHandler *handler.TagHandler,
	snsHandler *handler.SnsHandler,
	salesSiteHandler *handler.SalesSiteHandler,
	skillMarketHandler *handler.SkillMarketHandler,
	contactHandler *handler.ContactHandler,
	userHandler *handler.UserHandler,
	auth *middleware.AuthMiddleware,
	admin *middleware.AdminMiddleware,
	logging func(http.Handler) http.Handler,
	cors func(http.Handler) http.Handler,
) http.Handler {
	r := mux.NewRouter().StrictSlash(true)
	r.Use(logging)
	requireAdmin := func(hf http.HandlerFunc) http.Handler {
		return auth.RequireSession(admin.RequireAdmin(http.HandlerFunc(hf)))
	}

	// health check
	r.HandleFunc("/api/health_check", healthHandler.Check).Methods(http.MethodGet)

	// category
	r.HandleFunc("/api/category", categoryHandler.List).Methods(http.MethodGet)
	r.Handle("/api/category", requireAdmin(categoryHandler.Create)).Methods(http.MethodPost)
	r.Handle("/api/category/{category_uuid}", requireAdmin(categoryHandler.Update)).Methods(http.MethodPut)
	r.Handle("/api/category/{category_uuid}", requireAdmin(categoryHandler.Delete)).Methods(http.MethodDelete)

	// target
	r.HandleFunc("/api/target", targetHandler.List).Methods(http.MethodGet)
	r.Handle("/api/target", requireAdmin(targetHandler.Create)).Methods(http.MethodPost)
	r.Handle("/api/target/{target_uuid}", requireAdmin(targetHandler.Update)).Methods(http.MethodPut)
	r.Handle("/api/target/{target_uuid}", requireAdmin(targetHandler.Delete)).Methods(http.MethodDelete)

	// tag
	r.HandleFunc("/api/tag", tagHandler.List).Methods(http.MethodGet)
	r.Handle("/api/tag", requireAdmin(tagHandler.Create)).Methods(http.MethodPost)
	r.Handle("/api/tag/{tag_uuid}", requireAdmin(tagHandler.Update)).Methods(http.MethodPut)
	r.Handle("/api/tag/{tag_uuid}", requireAdmin(tagHandler.Delete)).Methods(http.MethodDelete)

	// sns
	r.HandleFunc("/api/sns", snsHandler.List).Methods(http.MethodGet)
	r.Handle("/api/sns", requireAdmin(snsHandler.Create)).Methods(http.MethodPost)
	r.Handle("/api/sns/{sns_uuid}", requireAdmin(snsHandler.Update)).Methods(http.MethodPut)
	r.Handle("/api/sns/{sns_uuid}", requireAdmin(snsHandler.Delete)).Methods(http.MethodDelete)

	// sales site
	r.HandleFunc("/api/sales_site", salesSiteHandler.List).Methods(http.MethodGet)
	r.Handle("/api/sales_site", requireAdmin(salesSiteHandler.Create)).Methods(http.MethodPost)
	r.Handle("/api/sales_site/{sales_site_uuid}", requireAdmin(salesSiteHandler.Update)).Methods(http.MethodPut)
	r.Handle("/api/sales_site/{sales_site_uuid}", requireAdmin(salesSiteHandler.Delete)).Methods(http.MethodDelete)

	// skill market
	r.HandleFunc("/api/skill_market", skillMarketHandler.List).Methods(http.MethodGet)
	r.Handle("/api/skill_market", requireAdmin(skillMarketHandler.Create)).Methods(http.MethodPost)
	r.Handle("/api/skill_market/{skill_market_uuid}", requireAdmin(skillMarketHandler.Update)).Methods(http.MethodPut)
	r.Handle("/api/skill_market/{skill_market_uuid}", requireAdmin(skillMarketHandler.Delete)).Methods(http.MethodDelete)

	// contact
	r.Handle("/api/contact", requireAdmin(contactHandler.List)).Methods(http.MethodGet)
	r.HandleFunc("/api/contact", contactHandler.Create).Methods(http.MethodPost)

	// user
	r.Handle("/api/user/login", auth.RequireSession(http.HandlerFunc(userHandler.GetLoginUser))).Methods(http.MethodGet)
	r.HandleFunc("/api/user/login", userHandler.Login).Methods(http.MethodPost)
	r.Handle("/api/user/logout", auth.RequireSession(http.HandlerFunc(userHandler.Logout))).Methods(http.MethodPost)

	return cors(r)
}
