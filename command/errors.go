package command

import "errors"

var (
	ErrCommandPrefixNotMatch = errors.New("command prefix mismatched")
	ErrInvalidPrimaryArg     = errors.New("invalid primary arg")
	ErrPrimaryArgExist       = errors.New("primary arg already exists")
	ErrCommandPrefixConflict = errors.New("primary arg conflict with command prefix")
	ErrCommandNotFound       = errors.New("command not found")
)
