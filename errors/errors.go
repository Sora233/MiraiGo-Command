package errors

import "errors"

var (
	ErrNotInitialized = errors.New("not initlized")
	ErrParseError = errors.New("command parser error")
)
