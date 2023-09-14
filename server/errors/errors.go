package errors

import "fmt"

type AppError struct {
	Code    int    // Custom error code
	Message string // Error message
}

func (e *AppError) Error() string {
	return e.Message
}

func Wrap(err error, message string) error {
	return &AppError{
		Code:    500,
		Message: fmt.Sprintf("%s: %v", message, err),
	}
}
