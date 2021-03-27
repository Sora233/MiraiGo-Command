package runtime

import (
	"github.com/Mrs4s/MiraiGo/client"
)

type Context struct {
	*Parser
	*client.QQClient
}

func NewContext(qqClient *client.QQClient, p *Parser) *Context {
	return &Context{
		p,
		qqClient,
	}
}
