package error

type NotFoundError struct {
	message string
}
type BadRequestError struct {
	message string
}

func (e *NotFoundError) Error() string {
	return e.message
}

func (e *BadRequestError) Error() string {
	return e.message
}

func NewNotFoundError(s string) *NotFoundError {
	return &NotFoundError{
		message: s,
	}
}

func NewBadRequestError(s string) *BadRequestError {
	return &BadRequestError{
		message: s,
	}
}
