package libeloula

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/syseditor/libeloula/utils"
	"github.com/syseditor/libeloula/commands"
)

var commandMap = make(map[string]cmd.Command, 0)

func Initialize() {
	utils.LoadOps()
	RegisterCommands()
}

func RegisterCommands() {
	commandMap["op"] = cmd.New("op", "Gives OP permissions to a specific player", []string{}, commands.Op{})
}

func GetCommandMap() *map[string]cmd.Command {
	return &commandMap
}