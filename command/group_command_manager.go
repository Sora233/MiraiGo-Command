package command

import (
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"strings"
)

type LspGroupCommandManager struct {
	commandMap    map[string]Handler
	commandPrefix string
}

func NewLspGroupCommandManager(commandPrefix string) *LspGroupCommandManager {
	return &LspGroupCommandManager{
		commandMap:    make(map[string]Handler),
		commandPrefix: commandPrefix,
	}
}

func (m *LspGroupCommandManager) Execute(qqClient *client.QQClient, msg *message.GroupMessage) error {
	p := NewParserFromMessage(msg.Elements)
	if m.commandPrefix != "" && !strings.HasPrefix(p.Command, m.commandPrefix) {
		return ErrCommandPrefixNotMatch
	}
	h := m.commandMap[p.Command[1:]]
	if h == nil {
		return ErrCommandNotFound
	}
	h(NewContext(qqClient, p), msg)
	return nil
}

func (m *LspGroupCommandManager) Register(primaryArg string, h Handler) error {
	primaryArg = strings.TrimSpace(primaryArg)
	if !m.checkPrimaryArg(primaryArg) {
		return ErrInvalidPrimaryArg
	}
	if _, found := m.commandMap[strings.TrimSpace(primaryArg)]; found {
		return ErrPrimaryArgExist
	}
	m.commandMap[primaryArg] = h
	return nil
}

func (m *LspGroupCommandManager) checkPrimaryArg(arg string) bool {
	if strings.HasPrefix(arg, m.commandPrefix) {
		return false
	}
	return true
}
