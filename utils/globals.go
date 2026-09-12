package utils

import (
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
)

var Server *server.Server
var CommandMap map[string]cmd.Command
