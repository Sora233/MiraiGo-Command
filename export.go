package command

import (
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/Sora233/MiraiGo-Command/runtime"
)

var manager *Manager = NewManager("/")

func StdManager() *Manager {
	return manager
}

func SetPrefix(prefix string) {
	manager.commandPrefix = prefix
}

func ExecutePrivate(qqClient *client.QQClient, msg *message.PrivateMessage) error {
	return manager.ExecutePrivate(qqClient, msg)
}

func ExecuteGroup(qqClient *client.QQClient, msg *message.GroupMessage) error {
	return manager.ExecuteGroup(qqClient, msg)
}

func RegisterGroupHandler(primaryArg string, h GroupHandler) error {
	return manager.RegisterGroupHandler(primaryArg, h)
}

func RegisterPrivateHandler(primaryArg string, h PrivateHandler) error {
	return manager.RegisterPrivateHandler(primaryArg, h)
}

func DelegateGroup(handler ...runtime.ErrHandler) MiraiGoGroupHandler {
	return manager.DelegateGroup(handler...)
}

func DelegatePrivate(handler ...runtime.ErrHandler) MiraiGoPrivateHandler {
	return manager.DelegatePrivate(handler...)
}
