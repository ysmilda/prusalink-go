package printer

import "errors"

var (
	ErrNotFound             = errors.New("not found")
	ErrUnauthorized         = errors.New("unauthorized")
	ErrBadRequest           = errors.New("bad request")
	ErrNotModified          = errors.New("not modified")
	ErrForbidden            = errors.New("forbidden")
	ErrRequestTimeout       = errors.New("request timeout")
	ErrUnsupportedMediaType = errors.New("unsupported media type")
	ErrInternalServerError  = errors.New("internal server error")
	ErrNotImplemented       = errors.New("not implemented")
	ErrServiceUnavailable   = errors.New("service unavailable")
	ErrNonGcodeFile         = errors.New("non-GCode file")
	ErrEmptyFile            = errors.New("empty file")
	ErrStorageNotFound      = errors.New("storage not found")
	ErrConflict             = errors.New("file already in use")
	ErrEmptyBody            = errors.New("empty body")
)
