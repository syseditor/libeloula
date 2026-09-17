package commands

import (
	"fmt"
	"strconv"
	"time"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/syseditor/libeloula/form"
	"github.com/syseditor/libeloula/memcached"
	"github.com/syseditor/libeloula/utils"
)

type Profile struct {
	Player cmd.Optional[string] `cmd:"profile"`
}

func (c Profile) Run(src cmd.Source, output *cmd.Output, tx *world.Tx) {
	profilePlayer, _ := c.Player.Load()
	session, err := memcached.CacheManager.Load(profilePlayer)

	if err != nil {
		if pl, ok := src.(*player.Player); ok {
			pl.SendForm(form.NewProfileForm(*session))
		} else {
			timestampInt, err := strconv.ParseInt(fmt.Sprintf("%d", session.JoinedAt), 10, 64)
			utils.Check(err)

			timestamp := time.Unix(timestampInt, 0)

			output.Printf(
				"Profile\n\nUUID: %s\nUsername: %s\nFirst joined at: %s\nTotal blocks broken: %d\n",
				session.UUID,
				session.Username,
				timestamp.String(),
				session.BlocksBroken,
			)
		}
	} else {
		output.Errorf("%sPlayer %s not found in cache.", text.Red, profilePlayer)
	}
}
