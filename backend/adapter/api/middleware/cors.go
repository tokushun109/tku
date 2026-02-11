package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

func NewCORS(allowedOrigins []string, allowCredentials bool) func(http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowedOrigins:   allowedOrigins,
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: allowCredentials,
	})
	return c.Handler
}
