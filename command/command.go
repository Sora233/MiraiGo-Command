package command

import (
	"github.com/Logiase/MiraiGo-Template/bot"
	"github.com/Sora233/MiraiGo-Command/errors"
)

func Init() {
	if commandManager == nil {
		commandManager = new(SimpleCommandManager)
	}
	bot.RegisterModule(commandManager)
}

func SetCommandManager(cm CommandManager) {
	commandManager = cm
}

func GetCommandManager() (CommandManager, error) {
	if commandManager == nil {
		return nil, errors.ErrNotInitialized
	}
	return commandManager, nil
}

type CommandScope int64

const (
	PrivateMessage CommandScope = 1 << iota
	GroupMessage
)

type Command interface {
	CommandFormat
	CommandScope() CommandScope
	Register() error
}

type SimpleCommand struct {
	SimpleCommandFormat
	Scope CommandScope
}

func (sc *SimpleCommand) CommandScope() CommandScope {
	return sc.Scope
}

func (sc *SimpleCommand) Register() error {
	return nil
}

func (sc *SimpleCommand) AddScope(scopes ...CommandScope) error {
	for _, scope := range scopes {
		sc.Scope |= scope
	}
	return nil
}

func (sc *SimpleCommand) DeleteScope(scopes ...CommandScope) error {
	for _, scope := range scopes {
		sc.Scope ^= sc.Scope & scope
	}
	return nil
}
