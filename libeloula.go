package libeloula

import (
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/syseditor/libeloula/commands"
	"github.com/syseditor/libeloula/source/console"
	"github.com/syseditor/libeloula/utils"
)

func Initialize(instance *server.Server) {
	utils.Server = instance

	utils.LoadOps()
	createCommands()
	registerAllCommands()
}

func createCommands() {
	utils.CommandMap["op"] = cmd.New("op", "Gives OP permissions to a specific player", []string{"op"}, commands.Op{})
}

func registerAllCommands() {
	for _, cm := range utils.CommandMap {
		cmd.Register(cm)
	}
}

// Might remove
func GetCommandMap() *map[string]cmd.Command {
	return &utils.CommandMap
}

func StartConcoleBuffer() {
	console.InitBuffer()
}
