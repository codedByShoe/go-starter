package errors

import "errors"

var (
	ErrNotFound    = errors.New("resource not found")
	ErrBadRequest  = errors.New("bad request")
	ErrServerError = errors.New("internal server error")
)

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
