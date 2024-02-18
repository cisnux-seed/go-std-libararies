package utils

type ValidationError struct {
	message string
}

func New(message string) *ValidationError {
	return &ValidationError{message: message}
}

func (err *ValidationError) Error() (message string) {
	message = err.message
	return
}
