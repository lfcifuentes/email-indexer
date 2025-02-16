package coreerrors

// Error types

// UsernamePathNotFound is an error type for when the username path is not found
// in the request
type UsernamePathNotFound struct {
	Message string
}

// MissingRequiredParameter is an error type for when a required parameter is
// missing in the request
type MissingRequiredParameter struct {
	Message string
}
