package utils

type ErrorType int16

const (
	Validation ErrorType = iota
	InternalError
	BusinessRuleViolation
)

type UseCaseError struct {
	Type    ErrorType
	Message string
}

func (e *UseCaseError) Error() string {
	return e.Message
}
