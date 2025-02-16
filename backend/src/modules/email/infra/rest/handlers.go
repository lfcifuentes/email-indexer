package emailrest

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
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

// RefreshByUserHandler
//
//	@Summary	Refresh all emails for a given user
//	@Schemes
//	@Description	Refresh all emails for a given user
//	@Tags			Emails
//	@Accept			json
//	@Produce		json
//	@Param			username	path		string	true	"Username"
//	@Success		200			{object}	corehttp.ResponseModel
//	@Failure		400			{object}	corehttp.ErrorResponse
//	@Router			/refresh/{username} [get]
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

// RefreshAllEmailHandler
//
//	@Summary	Refresh all emails for all users
//	@Schemes
//	@Description	Refresh all emails for all users
//	@Tags			Emails
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	corehttp.ResponseModel
//	@Failure		400	{object}	corehttp.ErrorResponse
//	@Router			/refresh [get]
func (h Handler) RefreshAllEmailHandler(w http.ResponseWriter, r *http.Request) {
	err := h.useCases.RefreshAllEmails()
	if err != nil {
		corehttp.SendError(w, r, err)
		return
	}

	corehttp.SendSuccess(w, r, "Emails refreshed successfully", nil)
}

// GetUserNamesHandler returns all the usernames
//
//	@Summary	Get all usernames
//	@Schemes
//	@Description	Get all usernames
//	@Tags			Emails
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	corehttp.ResponseModel
//	@Failure		400	{object}	corehttp.ErrorResponse
//	@Router			/usernames [get]
func (h Handler) GetUserNamesHandler(w http.ResponseWriter, r *http.Request) {

	users, err := h.useCases.GetUserNames()
	if err != nil {
		corehttp.SendError(w, r, err)
		return
	}

	corehttp.SendSuccess(w, r, "Users found successfully", users)
}
