package runtime

import "errors"

type ErrHandler func(err error)

func DefaultErrHandler(err error) {

}

var (
	ErrCommandPrefixNotMatch = errors.New("command prefix mismatched")
	ErrInvalidPrimaryArg     = errors.New("invalid primary arg")
	ErrPrimaryArgExist       = errors.New("primary arg already exists")
	ErrCommandPrefixConflict = errors.New("primary arg conflict with command prefix")
	ErrCommandNotFound       = errors.New("command not found")
	ErrInvalidHandler        = errors.New("invalid handler")
)
