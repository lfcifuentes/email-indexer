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

func NewUsernamePathNotFoundError(err error, path string) error {
	msg := fmt.Sprintf("username path not found: %s", path)
	return UsernamePathNotFound{
		buildMessage(nil, &msg),
	}
}

func (e UsernamePathNotFound) Error() string {
	return e.Message
}

func NewMissingRequiredParameterError(param string) error {
	msg := fmt.Sprintf("missing required parameter: %s", param)
	return MissingRequiredParameter{
		buildMessage(nil, &msg),
	}
}

func (e MissingRequiredParameter) Error() string {
	return e.Message
}
