package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/lfcifuentes/email-indexer/backend/src/http/middleware"
	emailrest "github.com/lfcifuentes/email-indexer/backend/src/modules/email/infra/rest"
	searchrest "github.com/lfcifuentes/email-indexer/backend/src/modules/search/infra/rest"
	"github.com/lfcifuentes/email-indexer/backend/src/services/email"
	"github.com/lfcifuentes/email-indexer/backend/src/services/zinc"
)

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
