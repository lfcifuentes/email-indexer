package searchrest

import (
	"net/http"

	coreerrors "github.com/lfcifuentes/email-indexer/backend/src/core/errors"
	corehttp "github.com/lfcifuentes/email-indexer/backend/src/core/http"
	searchapp "github.com/lfcifuentes/email-indexer/backend/src/modules/search/app"
)

type Handler struct {
	useCases *searchapp.SearchUseCases
}

func NewHandler(useCases *searchapp.SearchUseCases) *Handler {
	return &Handler{
		useCases: useCases,
	}
}

// SearchHandler returns all the emails that match the search criteria
//
//	@Summary	Search emails
//	@Schemes
//	@Description	Search emails by query string
//	@Tags			Search
//	@Accept			json
//	@Produce		json
//	@Param			query	query		string	true	"Search query"
//	@Param			page	query		string	false	"Page number"
//	@Success		200		{object}	corehttp.ResponseModel
//	@Failure		400		{object}	corehttp.ErrorResponse
//	@Router			/search [get]
func (h Handler) SearchHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve search criteria from the URL query parameters.
	query := r.URL.Query().Get("query")
	if query == "" {
		err := coreerrors.NewMissingRequiredParameterError("query")
		corehttp.SendError(w, r, err)
		return
	}
	// Retrieve search criteria from the URL query parameters.
	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}
	emails, err := h.useCases.SearchEmails(query, page)
	if err != nil {
		corehttp.SendError(w, r, err)
		return
	}
	corehttp.SendSuccess(w, r, "Search results", emails)
}
