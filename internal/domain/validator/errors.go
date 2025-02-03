package validator

import "fmt"

type ValidationError struct {
	Code    int
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func NewValidationError(code int, message string) *ValidationError {
	return &ValidationError{
		Code:    code,
		Message: message,
	}
}
