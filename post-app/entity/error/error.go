package error

type DomainError struct {
	Message string
}

func (e *DomainError) Error() string {
	return e.Message
}

var (
	SystemError  = &DomainError{Message: "system error"}
	NotFound     = &DomainError{Message: "not found"}
	InvalidInput = &DomainError{Message: "invalid input"}
)
