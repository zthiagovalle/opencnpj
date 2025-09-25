package opencnpj

import "errors"

var (
	// ErrNotFound is returned when a CNPJ is not found.
	ErrNotFound        = errors.New("not found")
	// ErrTooManyRequests is returned when the API rate limit is exceeded.
	ErrTooManyRequests = errors.New("too many requests")
)
