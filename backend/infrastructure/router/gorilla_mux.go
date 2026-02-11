package router

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tokushun109/tku/backend/adapter/api/action"
	"github.com/tokushun109/tku/backend/adapter/api/middleware"
	"github.com/tokushun109/tku/backend/adapter/logger"

	"github.com/gorilla/mux"
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

	clientURL := os.Getenv("CLIENT_URL")
	if clientURL == "" {
		panic("CLIENT_URL is required for CORS configuration")
	}

	var handler http.Handler = r
	handler = middleware.NewRecovery(log)(handler)
	handler = middleware.NewLogger(log)(handler)
	handler = middleware.NewCORS([]string{clientURL}, true)(handler)

	srv := &GorillaMuxServer{
		router:  r,
		handler: handler,
		port:    port,
		db:      db,
		log:     log,
	}

	srv.registerRoutes(r)
	return srv
}

func (g *GorillaMuxServer) registerRoutes(r *mux.Router) {
	g.registerHealthRoutes(r)
}

func (g *GorillaMuxServer) registerHealthRoutes(r *mux.Router) {
	healthCheck := action.NewHealthCheckAction(g.db, g.log)
	r.HandleFunc("/api/health_check", healthCheck.Execute).Methods(http.MethodGet)
}

func (g *GorillaMuxServer) Listen() {
	server := &http.Server{
		Addr:         ":" + string(g.port),
		Handler:      g.handler,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			if g.log != nil {
				g.log.Errorf("server error: %v", err)
			}
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		if g.log != nil {
			g.log.Errorf("shutdown error: %v", err)
		}
	}
}
