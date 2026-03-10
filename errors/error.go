package errors

// Error represents a structured error used across the transport library.
//
// It provides a stable error code and human-readable message while
// remaining compatible with the standard error interface.
type Error struct {
	// Code identifies the error condition.
	Code string

	// Message describes the error.
	Message string
}

// Error implements the error interface.
func (e *Error) Error() string {
	if e.Message == "" {
		return e.Code
	}
	return e.Code + ": " + e.Message
}
