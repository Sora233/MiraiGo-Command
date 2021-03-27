package command

import (
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/Sora233/MiraiGo-Command/runtime"
	"strings"
)

type MiraiGoGroupHandler func(qqClient *client.QQClient, groupMessage *message.GroupMessage)
type MiraiGoPrivateHandler func(qqClient *client.QQClient, groupMessage *message.PrivateMessage)

type FuncType int

const (
	PrivateHandlerType FuncType = iota
	GroupHandlerType
)

type Func interface {
	Type() FuncType
}

type PrivateHandler func(ctx *runtime.Context, groupMsg *message.PrivateMessage)

func (p PrivateHandler) Type() FuncType {
	return PrivateHandlerType
}

type GroupHandler func(ctx *runtime.Context, groupMsg *message.GroupMessage)

func (g GroupHandler) Type() FuncType {
	return PrivateHandlerType
}

type Manager struct {
	commandMap    map[FuncType]map[string]Func
	commandPrefix string
}

func NewManager(commandPrefix string) *Manager {
	m := &Manager{
		commandMap:    make(map[FuncType]map[string]Func),
		commandPrefix: commandPrefix,
	}
	m.commandMap[PrivateHandlerType] = make(map[string]Func)
	m.commandMap[GroupHandlerType] = make(map[string]Func)
	return m
}

func (m *Manager) SetPrefix(prefix string) {
	m.commandPrefix = prefix
}

func (m *Manager) ExecutePrivate(qqClient *client.QQClient, msg *message.PrivateMessage) error {
	p := runtime.NewParserFromMessage(msg.Elements)
	if m.commandPrefix != "" && !strings.HasPrefix(p.Command, m.commandPrefix) {
		return runtime.ErrCommandPrefixNotMatch
	}
	h := m.commandMap[PrivateHandlerType][strings.TrimPrefix(p.Command, m.commandPrefix)]
	if h == nil {
		return runtime.ErrCommandNotFound
	}
	h.(PrivateHandler)(runtime.NewContext(qqClient, p), msg)
	return nil
}

func (m *Manager) ExecuteGroup(qqClient *client.QQClient, msg *message.GroupMessage) error {
	p := runtime.NewParserFromMessage(msg.Elements)
	if m.commandPrefix != "" && !strings.HasPrefix(p.Command, m.commandPrefix) {
		return runtime.ErrCommandPrefixNotMatch
	}
	h := m.commandMap[GroupHandlerType][strings.TrimPrefix(p.Command, m.commandPrefix)]
	if h == nil {
		return runtime.ErrCommandNotFound
	}
	h.(GroupHandler)(runtime.NewContext(qqClient, p), msg)
	return nil
}

func (m *Manager) register(primaryArg string, h Func) error {
	if h == nil {
		return runtime.ErrInvalidHandler
	}
	primaryArg = strings.TrimSpace(primaryArg)
	if _, found := m.commandMap[h.Type()][strings.TrimSpace(primaryArg)]; found {
		return runtime.ErrPrimaryArgExist
	}
	m.commandMap[h.Type()][primaryArg] = h
	return nil
}

func (m *Manager) RegisterGroupHandler(primaryArg string, h GroupHandler) error {
	return m.register(primaryArg, h)
}

func (m *Manager) RegisterPrivateHandler(primaryArg string, h PrivateHandler) error {
	return m.register(primaryArg, h)
}

func (m *Manager) DelegateGroup(handler ...runtime.ErrHandler) MiraiGoGroupHandler {
	return func(qqClient *client.QQClient, groupMessage *message.GroupMessage) {
		err := m.ExecuteGroup(qqClient, groupMessage)
		var h = runtime.DefaultErrHandler
		if len(handler) >= 1 {
			h = handler[0]
		}
		if h != nil {
			h(err)
		}
	}
}

func (m *Manager) DelegatePrivate(handler ...runtime.ErrHandler) MiraiGoPrivateHandler {
	return func(qqClient *client.QQClient, privateMessage *message.PrivateMessage) {
		err := m.ExecutePrivate(qqClient, privateMessage)
		var h = runtime.DefaultErrHandler
		if len(handler) >= 1 {
			h = handler[0]
		}
		if h != nil {
			h(err)
		}
	}
}
