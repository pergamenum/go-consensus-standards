package ehandler

import (
	"fmt"
	"strings"
)

// Wrap embeds the error e with additional information s.
func Wrap(s string, e error) error {

	clean := strings.TrimSpace(s)

	return fmt.Errorf("[%s] -> %w", clean, e)
}
