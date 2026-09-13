package console

import (
	"fmt"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

type ConsoleCommandSource struct{}

func (ConsoleCommandSource) World() *world.World {
	return nil // The Console Command Sender is not registered to any worlds
}

func (ConsoleCommandSource) Name() string {
	return "console"
}

func (ConsoleCommandSource) Position() mgl64.Vec3 {
	return mgl64.Vec3{} // (0,0,0)
}

func (ConsoleCommandSource) SendCommandOutput(o *cmd.Output) {
	for _, msg := range o.Messages() {
		fmt.Println(text.ANSI(msg.String()))
	}

	for _, err := range o.Errors() {
		fmt.Println(text.ANSI(err.Error()))
	}
}
