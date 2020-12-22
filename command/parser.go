package command

import "github.com/Sora233/MiraiGo-Command/errors"

type CommandParser interface {
	Parse(interface{}) (CommandFormat, error)
}

type SimpleCommandParser struct {
}

func (scp *SimpleCommandParser) Parse(i interface{}) (CommandFormat, error) {
	switch i.(type) {
	case string:
		return scp.ParseString(i.(string))

	default:
		return nil, errors.ErrParseError
	}
}

func (scp *SimpleCommandParser) ParseString(s string) (*SimpleCommandFormat, error) {
	return nil, nil
}
