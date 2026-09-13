package utils

import (
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/syseditor/libeloula/memcached"
)

var Server *server.Server
var CommandMap = make(map[string]cmd.Command, 0) // Might remove
var SessionManager *memcached.SessionManager
