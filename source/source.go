package source

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
)

type CommandSource interface {
	cmd.Source

	Name() string
	World() *world.World
}