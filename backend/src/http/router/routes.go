package router

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	chi_middleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	_ "github.com/lfcifuentes/email-indexer/backend/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// SetRoutes sets the routes for the router
//
// Parameters:
//
//	r: The chi router
//
// Returns:
//
//	None
//
// Usage:
//
//	r := chi.NewRouter()
//	SetRoutes(r)
func SetBaseRoutes(r *chi.Mux) {
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]string{"message": "pong"})
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]string{"message": "Welcome to the email indexer api!"})
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		reqID := chi_middleware.GetReqID(r.Context())
		// return a 200 status and request id
		render.JSON(w, r, map[string]string{"request_id": reqID})
		render.Status(r, http.StatusOK)
	})

	if os.Getenv("APP_ENV") == "development" {
		// Swagger
		r.Mount("/swagger", httpSwagger.WrapHandler)
	}
}
