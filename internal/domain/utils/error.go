package utils

type ErrorType int16

const (
	ValidationError ErrorType = iota
	InternalError
	BusinessRuleViolationError
)

type UseCaseError struct {
	Type    ErrorType
	Message string
}

func (e *UseCaseError) Error() string {
	return e.Message
}
