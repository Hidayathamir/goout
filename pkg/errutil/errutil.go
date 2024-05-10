package errutil

import (
	"fmt"

	"github.com/Hidayathamir/goout/pkg/runtime"
)

// WrapOpt represents options for the Wrap function.
type WrapOpt struct {
	Skip int
}

// WrapOption represents an option for the Wrap function.
type WrapOption func(*WrapOpt)

const defaultSkip = 2

// Wrap wraps the given error with the name of the calling function.
func Wrap(err error, options ...WrapOption) error {
	option := &WrapOpt{
		Skip: defaultSkip,
	}
	for _, opt := range options {
		opt(option)
	}
	return fmt.Errorf(runtime.FuncName(runtime.WithSkip(option.Skip))+": %w", err)
}

// WithSkip sets the number of stack frames to skip when identifying the caller.
func WithSkip(skip int) WrapOption {
	return func(wo *WrapOpt) {
		wo.Skip = skip + defaultSkip
	}
}
