package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/lfcifuentes/email-indexer/backend/src/http/middleware"
	emailrest "github.com/lfcifuentes/email-indexer/backend/src/modules/email/infra/rest"
	searchrest "github.com/lfcifuentes/email-indexer/backend/src/modules/search/infra/rest"
	"github.com/lfcifuentes/email-indexer/backend/src/services/email"
	"github.com/lfcifuentes/email-indexer/backend/src/services/zinc"
)

// NewRouter creates a new router
// It sets the base routes and the email and search routes
// It returns the chi router
//
// Parameters:
//
//	port: The port to run the server on
//	zincService: The zinc service
//	emailServices: The email services
//
// Returns:
//
//	*chi.Mux: The chi router
func NewRouter(
	port string,
	zincService *zinc.ZincService,
	emailServices *email.EmailServices,
) *chi.Mux {
	// Create a new chi router
	r := chi.NewRouter()

	// API base middlewares
	middleware.StartAllMiddlewares(r)

	// Base Routes
	SetBaseRoutes(r)

	// Email Routes
	emailrest.Router(r, emailServices, zincService)
	// Search Routes
	searchrest.Router(r, emailServices, zincService)

	return r
}
