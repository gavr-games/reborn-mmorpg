package game_test

import (
	"sync"

	"github.com/alicebob/miniredis/v2"
	"github.com/gavr-games/reborn-mmorpg/pkg/game"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/redis/go-redis/v9"
)

var once sync.Once
var e *game.Engine

func Setup() {
	once.Do(func() {
		server, _ := miniredis.Run()
		rdb := redis.NewClient(&redis.Options{
			Addr: server.Addr(),
		})
		storage.SetClient(rdb)
		e = game.NewEngine()
		e.Init(true)
	})
}

func GetEngine() *game.Engine {
	return e
}
