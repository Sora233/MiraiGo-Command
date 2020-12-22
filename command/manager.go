package command

import (
	"github.com/Logiase/MiraiGo-Template/bot"
	"sync"
)

type CommandManager interface {
	ExecuteCommandOn(CommandScope, Command)
	GetCommands() []Command

	bot.Module
}

var commandManager CommandManager

type SimpleCommandManager struct {
	commands []Command
}

func (scm *SimpleCommandManager) GetCommands() []Command {
	panic("implement me")
}

func (scm *SimpleCommandManager) ExecuteCommandOn(scope CommandScope, command Command) {
	panic("implement me")
}

func (c *SimpleCommandManager) MiraiGoModule() bot.ModuleInfo {
	return bot.ModuleInfo{
		ID:       "sora233.miraigo-commandManager",
		Instance: commandManager,
	}
}

func (c *SimpleCommandManager) Init() {
}

func (c *SimpleCommandManager) PostInit() {
}

func (c *SimpleCommandManager) Serve(bot *bot.Bot) {
	panic("implement me")
}

func (c *SimpleCommandManager) Start(bot *bot.Bot) {
}

func (c *SimpleCommandManager) Stop(bot *bot.Bot, wg *sync.WaitGroup) {
	defer wg.Done()
}
