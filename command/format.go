package command

type CommandFormat interface {
	Usage() string
	PrimaryArgs() CommandArgs
	Args() []CommandArgs
}

type SimpleCommandFormat struct {
	PrimaryName string
}


func (dcf *SimpleCommandFormat) Usage() string {
	return ""
}

func (dcf *SimpleCommandFormat) PrimaryArgs() CommandArgs {
	return nil

}

func (dcf *SimpleCommandFormat) Args() []CommandArgs {
	return nil
}


func NewSimpleCommandFromString(s string) (*SimpleCommandFormat, error) {
	parser := new(SimpleCommandParser)
	return parser.ParseString(s)
}
