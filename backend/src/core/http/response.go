package corehttp

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	coreerrors "github.com/lfcifuentes/email-indexer/backend/src/core/errors"
)

// ResponseModel represents a generic response model that includes
// a message and optional data.
// Fields:
//
//	Message: The message returned in the response.
//	Data: Optional data associated with the response, can be any type.
type ResponseModel struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// ErrorResponse represents a generic error response with a message.
// It includes a single field:
//
//	Message: The error message returned in the response.
type ErrorResponse struct {
	Message string `json:"message"`
}

// EmptyResponse represents a success response.
// It is an empty struct, meaning no data is returned in the response.
type EmptyResponse struct {
}

// ResponseSuccess returns a ResponseModel with the provided message and data.
func ResponseSuccess(message string, data interface{}) ResponseModel {
	return ResponseModel{
		Message: message,
		Data:    data,
	}
}

// ResponseEmpty returns an empty ResponseModel.
func ResponseEmpty() ResponseModel {
	return ResponseModel{
		Message: "",
		Data:    nil,
	}
}

// ResponseErr returns a ResponseModel representing an error.
func ResponseErr(err interface{}) ResponseModel {
	return ResponseModel{
		Message: fmt.Sprintf("%v", err),
		Data:    nil,
	}
}

func SendSuccess(w http.ResponseWriter, r *http.Request, message string, data interface{}) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, ResponseSuccess(message, data))
}

// SendEmpty sends an empty JSON response with a 200 OK status code using chi render.
func SendEmpty(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, ResponseEmpty())
}

// SendError sends an appropriate JSON error response based on the type of error using chi render.
func SendError(w http.ResponseWriter, r *http.Request, err error) {
	switch {
	case errors.As(err, &coreerrors.UsernamePathNotFound{}):
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, ResponseErr(err.Error()))
		return
	case errors.As(err, &coreerrors.MissingRequiredParameter{}):
		// Assuming you have a MissingRequiredParameter error type defined
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ResponseErr(err.Error()))
		return
	default:
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, ResponseErr(err))
		return
	}
}
