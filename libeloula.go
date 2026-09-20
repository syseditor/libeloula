package libeloula

import (
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/syseditor/libeloula/commands"
	"github.com/syseditor/libeloula/db"
	"github.com/syseditor/libeloula/memcached"
	"github.com/syseditor/libeloula/source/console"
	"github.com/syseditor/libeloula/utils"
)

var DataProvider db.DataProvider

func Initialize(instance *server.Server, addr []string, database_info [3]string) {
	utils.Server = instance

	utils.LoadOps()
	createCommands()
	registerAllCommands()

	memcached.CacheManager = memcached.NewSessionManager(addr...)
	DataProvider = db.NewDataProvider(database_info[0])
	DataProvider.InitializeDB(database_info[1], database_info[2])
}

func createCommands() {
	utils.CommandMap["op"] = cmd.New("op", "Gives OP permissions to a specific player", []string{"op"}, commands.Op{})
	utils.CommandMap["profile"] = cmd.New("profile", "Displays the profile of a player!", []string{"profile"}, commands.Profile{})
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
