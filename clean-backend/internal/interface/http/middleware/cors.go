package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

func CORSMiddleware(allowedOrigins []string) func(http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowedOrigins:   allowedOrigins,
		AllowCredentials: true,
	})
	return c.Handler
}
