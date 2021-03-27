package runtime

import "github.com/Mrs4s/MiraiGo/message"

type CommandArgType int64

const (
	Text CommandArgType = 1 + iota
	Image
)

type CommandArgs interface {
	Type() CommandArgType
}

type TextCommandArgs struct {
	Text string
}

func (sca *TextCommandArgs) Type() CommandArgType {
	return Text
}

type ImageCommandArgs struct {
	Image *message.ImageElement
}

func (ica *ImageCommandArgs) Type() CommandArgType {
	return Image
}
