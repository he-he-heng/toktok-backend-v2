package domain

import "toktok-backend-v2/pkg/errors"

var (
	ErrInternalServerError = errors.New("internal server errors")
	ErrNotFound            = errors.New("item is not found")
	ErrConflict            = errors.New("item is conflict")
	ErrBadParam            = errors.New("given param is not valid")
)
