package emailrest

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	coreerrors "github.com/lfcifuentes/email-indexer/backend/src/core/errors"
	corehttp "github.com/lfcifuentes/email-indexer/backend/src/core/http"
	emailapp "github.com/lfcifuentes/email-indexer/backend/src/modules/email/app"
)

type Handler struct {
	useCases *emailapp.EmailUseCases
}

func NewHandler(useCases *emailapp.EmailUseCases) *Handler {
	return &Handler{
		useCases: useCases,
	}
}

func (h Handler) RefreshByUserHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve "username" from the URL path parameters.
	username := strings.TrimSpace(chi.URLParam(r, "username"))
	if username == "" {
		err := coreerrors.NewMissingRequiredParameterError("username")
		corehttp.SendError(w, r, err)
		return
	}
	totalImported, err := h.useCases.RefreshUserEmails(username)
	if err != nil {
		corehttp.SendError(w, r, err)
		return
	}
	message := fmt.Sprintf("%v emails imported for user %v", totalImported, username)
	corehttp.SendSuccess(w, r, message, nil)
}

func (h Handler) RefreshAllEmailHandler(w http.ResponseWriter, r *http.Request) {
	err := h.useCases.RefreshAllEmails()
	if err != nil {
		render.JSON(w, r, map[string]string{"message": "error"})
		return
	}

	render.JSON(w, r, map[string]string{"message": "Emails refreshed successfully"})
}

// GetUserNamesHandler returns all the usernames
func (h Handler) GetUserNamesHandler(w http.ResponseWriter, r *http.Request) {

	users, err := h.useCases.GetUserNames()
	if err != nil {
		render.JSON(w, r, map[string]string{"message": "error"})
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"message": "Users found successfully",
		"users":   users,
	})
}
