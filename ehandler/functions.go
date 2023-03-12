package ehandler

import (
	"fmt"
	"strings"
)

// Wrap embeds the error 'e' with additional information 'i'.
//
//	If 'i' is an error, then any type it had will be lost.
//	Returns 'e' if 'i' is not of type error or string.
func Wrap(i any, e error) error {

	var info string
	if s, ok := i.(string); ok {
		info = s
	}

	if err, ok := i.(error); ok {
		info = err.Error()
	}

	if info == "" {
		return e
	}

	clean := strings.TrimSpace(info)

	return fmt.Errorf("[%s] -> %w", clean, e)
}
