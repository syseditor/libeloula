package memcached

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/syseditor/libeloula/utils"
)

type PlayerHandler struct {
	player.NopHandler
}

func HandleJoin(player *player.Player) { //call this in main.go
	_, err := CacheManager.LoadOrCreate(player.UUID(), player.Name())
	utils.Check(err)
}

func (PlayerHandler) HandleQuit(player *player.Player) {
	/*
		Solutions for Player Quit:
		- You may do nothing. Each session has an expiration date, in that case you'll need a service
		to retrieve all expiring sessions and save them to db
		- You can force save the session to the db and then delete the cached session
		- Do both, force save everything to db and keep the cached session till expiration. In that case, we might need to decrease the expiration time.

		For testing purposes, we'll be deleting all sessions without saving them.
	*/

	err := CacheManager.Delete(player.UUID())
	utils.Check(err)
}

func (PlayerHandler) HandleBlockBreak(ctx *player.Context, _ cube.Pos, _ *[]item.Stack, _ *int) {
	CacheManager.AddBlocksBroken(ctx.Player().UUID())
}
