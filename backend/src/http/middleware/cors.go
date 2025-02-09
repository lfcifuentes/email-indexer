package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/cors"
)

// CorsMiddleware is a middleware function that sets up CORS (Cross-Origin Resource Sharing)
// for the HTTP server. It reads the allowed origins from the environment variable
// "ACCEPTED_DOMAINS", and if not set, defaults to allowing all origins ("*").
// The middleware allows the following HTTP methods: GET, POST, PUT, HEAD, OPTIONS, DELETE.
// It also allows the headers: X-Requested-With and Content-Type.
//
// Usage:
//
//	r := chi.NewRouter()
//	r.Use(middleware.CorsMiddleware)
func CorsMiddleware(next http.Handler) http.Handler {
	allowedOrigins := strings.Split(os.Getenv("ACCEPTED_DOMAINS"), ",")
	if len(allowedOrigins) == 0 || (len(allowedOrigins) == 1 && allowedOrigins[0] == "") {
		allowedOrigins = []string{"*"}
	}

	return cors.Handler(cors.Options{
		AllowedOrigins: allowedOrigins,
		AllowedMethods: []string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"},
		AllowedHeaders: []string{"X-Requested-With", "Content-Type"},
	})(next)
}
