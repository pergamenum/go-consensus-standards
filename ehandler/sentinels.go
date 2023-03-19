package ehandler

import (
	"errors"
)

// This is a sentinel error, intended for use with Wrap() and errors.Is().
var (
	ErrConflict   = errors.New("[CONFLICT]")
	ErrNotFound   = errors.New("[NOT FOUND]")
	ErrInternal   = errors.New("[INTERNAL ERROR]")
	ErrCorrupt    = errors.New("[CORRUPT STATE]")
	ErrBadRequest = errors.New("[BAD REQUEST]")
	ErrBadGateway = errors.New("[BAD GATEWAY]")
)
