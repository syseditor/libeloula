package libeloula

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/syseditor/libeloula/commands"
	"github.com/syseditor/libeloula/utils"
)

var commandMap = make(map[string]cmd.Command, 0)

func Initialize() {
	utils.LoadOps()
	createCommands()
	registerAllCommands()
}

func createCommands() {
	commandMap["op"] = cmd.New("op", "Gives OP permissions to a specific player", []string{"op"}, commands.Op{})
}

func registerAllCommands() {
	for _, cm := range commandMap {
		cmd.Register(cm)
	}
}

func GetCommandMap() *map[string]cmd.Command {
	return &commandMap
}
