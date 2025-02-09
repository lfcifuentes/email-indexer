package emailrest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	emailapp "github.com/lfcifuentes/email-indexer/backend/src/modules/email/app"
	"github.com/lfcifuentes/email-indexer/backend/src/services/email"
	"github.com/lfcifuentes/email-indexer/backend/src/services/zinc"
)

func Router(router *chi.Mux, emailServices *email.EmailServices, zincService *zinc.ZincService) {
	useCases := emailapp.NewUseCase(emailServices, zincService)
	handler := NewHandler(useCases)

	router.Route("/email", func(r chi.Router) {
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			render.JSON(w, r, map[string]string{"message": "email pong"})
		})
		// update email zinc data
		r.Get("/refresh", handler.RefreshAllEmailHandler)
		// update email zinc data
		r.Get("/refresh/{username}", handler.RefreshByUserHandler)
		// get all usernames for a given email
		r.Get("/usernames", handler.GetUserNamesHandler)
	})

}
