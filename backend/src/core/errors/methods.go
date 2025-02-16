package coreerrors

import "fmt"

// buildMessage constructs a detailed error message by combining
// the provided message with the error's message (if any).
// Parameters:
//
//	err: The error that occurred.
//	message: A pointer to the base message string.
//
// Returns:
//
//	A concatenated string with the base message and the error message.
func buildMessage(err error, message *string) string {
	msg := *message
	if err != nil {
		msg = msg + ": " + err.Error()
	}
	return msg
}

// NewUsernamePathNotFoundError creates a new UsernamePathNotFound error.
// Parameters:
//
// err: The error that occurred.
// path: The path that was not found.
//
// Returns:
// A new UsernamePathNotFound error.
//
// Example:
// err := NewUsernamePathNotFoundError(err, "path")
//
//	if err != nil {
//	   log.Fatal(err)
//	}
func NewUsernamePathNotFoundError(err error, path string) error {
	msg := fmt.Sprintf("username path not found: %s", path)
	return UsernamePathNotFound{
		buildMessage(nil, &msg),
	}
}

// Error returns the error message.
func (e UsernamePathNotFound) Error() string {
	return e.Message
}

// NewMissingRequiredParameterError creates a new MissingRequiredParameter error.
// Parameters:
// param: The name of the missing parameter.
// Returns:
// A new MissingRequiredParameter error.
func NewMissingRequiredParameterError(param string) error {
	msg := fmt.Sprintf("missing required parameter: %s", param)
	return MissingRequiredParameter{
		buildMessage(nil, &msg),
	}
}

// Error returns the error message.
func (e MissingRequiredParameter) Error() string {
	return e.Message
}
