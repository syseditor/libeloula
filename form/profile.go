package form

import (
	"fmt"
	"strconv"
	"time"

	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/form"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/syseditor/libeloula/memcached"
	"github.com/syseditor/libeloula/utils"
)

type ProfileForm struct{}

func NewProfileForm(session memcached.PlayerSession) form.Menu {
	timestampInt, err := strconv.ParseInt(fmt.Sprintf("%d", session.JoinedAt), 10, 64)
	utils.Check(err)

	timestamp := time.Unix(timestampInt, 0)
	return form.NewMenu(
		ProfileForm{}, "Profile",
	).WithBody(
		form.NewLabel("UUID:"+session.UUID),
		form.NewLabel("Username: "+session.Username),
		form.NewLabel("First joined at: "+timestamp.String()),
		form.NewLabel(fmt.Sprintf("Total blocks broken: %d", session.BlocksBroken)),
	).WithButtons(
		form.NewButton("Great!", ""),
	)
}

func (f ProfileForm) Submit(submitter form.Submitter, pressed form.Button, world *world.Tx) {
	submitter.CloseForm()
	pl, _ := submitter.(*player.Player)
	pl.Data().Session.SendMessage("Closed profile form!")
}
