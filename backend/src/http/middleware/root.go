package middleware

import (
	"time"

	"github.com/go-chi/chi/v5"
	chi_middleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

// StartAllMiddlewares initializes and attaches a set of middlewares to the provided router.
// The middlewares included are:
// - RequestID: Adds a unique request ID to each request.
// - Logger: Logs the start and end of each request with the elapsed processing time.
// - Recoverer: Recovers from panics and returns a 500 internal server error.
// - URLFormat: Parses the URL extension to set the response content type.
// - SetContentType: Sets the response content type to JSON.
// - CorsMiddleware: Adds custom CORS headers to each request.
// - Timeout: Sets a timeout value on the request context to 60 seconds.
func StartAllMiddlewares(r *chi.Mux) {
	r.Use(
		// add request id middleware
		chi_middleware.RequestID,
		// enalbe request logging
		chi_middleware.Logger,
		// enable app panic recovery
		chi_middleware.Recoverer,
		// enable url format
		chi_middleware.URLFormat,
		// set content type
		render.SetContentType(render.ContentTypeJSON),
		// add custom cors middleware
		CorsMiddleware,
		// Set a timeout value on the request context (ctx), that will signal
		chi_middleware.Timeout(60*time.Second),
	)
}
