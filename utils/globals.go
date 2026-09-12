package utils

import (
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/syseditor/libeloula/connection"
)

var Server *server.Server
var CommandMap map[string]cmd.Command
var GlobalConnectionManager *connection.ConnectionManager