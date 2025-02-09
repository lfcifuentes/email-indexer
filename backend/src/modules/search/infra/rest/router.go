package searchrest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	searchapp "github.com/lfcifuentes/email-indexer/backend/src/modules/search/app"
	"github.com/lfcifuentes/email-indexer/backend/src/services/email"
	"github.com/lfcifuentes/email-indexer/backend/src/services/zinc"
)

func Router(router *chi.Mux, emailServices *email.EmailServices, zincService *zinc.ZincService) {
	useCases := searchapp.NewUseCase(emailServices, zincService)
	handler := NewHandler(useCases)

	router.Route("/search", func(r chi.Router) {
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			render.JSON(w, r, map[string]string{"message": "search pong"})
		})
		// search emails by request filters
		r.Get("/", handler.SearchHandler)
		// get all emails for a given username
		//r.Get("/user/{id}", handler.SearchHandler)
	})

}
