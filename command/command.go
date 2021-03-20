package command

import (
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
)

type Context struct {
	*parser
	*client.QQClient
}

type Handler func(ctx *Context, groupMsg *message.GroupMessage)

func NewContext(qqClient *client.QQClient, p *parser) *Context {
	return &Context{
		p,
		qqClient,
	}
}
