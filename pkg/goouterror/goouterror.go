// Package goouterror provides custom error definitions.
package goouterror

import "errors"

// ErrInvalidRequest represents an invalid request error.
var ErrInvalidRequest = errors.New("invalid request")
