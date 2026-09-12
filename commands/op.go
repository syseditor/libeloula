package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/syseditor/libeloula/source/console"
	"github.com/syseditor/libeloula/utils"
)

type Op struct {
	Player string `cmd:"player"`
}

func (t Op) Run(src cmd.Source, output *cmd.Output, tx *world.Tx) {
	if _, ok := src.(*console.ConsoleCommandSource); ok {
		op(t.Player, output)
	} else if sender, ok := src.(*player.Player); ok {
		if utils.IsOp(sender.Name()) {
			op(t.Player, output)
			if playerHandle, online := utils.Server.PlayerByName(t.Player); online {
				utils.SendOperatorAbilityPacket(playerHandle, true)
			} else {
				output.Printf("%sThe player %s seems to be offline.", text.Grey, t.Player)
			}
		} else {
			output.Error(text.DarkRed, "You do not have permission to execute this command.")
		}
	}
}

func op(player string, output *cmd.Output) {
	if utils.AddOp(player) {
		output.Printf("Successfully opped player %s.\n", player)
	} else {
		output.Errorf("Player %s is already opped.", player)
	}
}

type Deop struct {
	Player string `cmd:"player"`
}

func (t Deop) Run(src cmd.Source, output *cmd.Output, tx *world.Tx) {

}

type OpList struct{}

func (t OpList) Run(src cmd.Source, output *cmd.Output, tx *world.Tx) {

}
