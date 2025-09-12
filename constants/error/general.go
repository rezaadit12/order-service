package error

import "errors"

const (
	Success = "success"
	Error   = "error"
)

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrSQLError            = errors.New("database server failed to execute query")
	ErrTooManyRequests     = errors.New("too many request")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrInvalidToken        = errors.New("invalid token")
	ErrInvalidUploadFile   = errors.New("invalid upload file")
	ErrForbidden           = errors.New("forbidden")
	ErrSizeTooBig          = errors.New("size too big")
)

var GeneralErrors = []error{
	ErrInternalServerError,
	ErrSQLError,
	ErrTooManyRequests,
	ErrUnauthorized,
	ErrInvalidToken,
	ErrForbidden,
}
