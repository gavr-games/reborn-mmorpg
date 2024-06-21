package game_test

import (
	"sync"

	"github.com/alicebob/miniredis/v2"
	"github.com/gavr-games/reborn-mmorpg/pkg/game"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/constants"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/redis/go-redis/v9"
)

var once sync.Once
var e *game.Engine
var surfaceArea *entity.GameArea

func Setup() {
	once.Do(func() {
		server, _ := miniredis.Run()
		rdb := redis.NewClient(&redis.Options{
			Addr: server.Addr(),
		})
		storage.SetClient(rdb)
		e = game.NewEngine()
		e.Init(true)
		e.EnableTestingMode()
		surfaceArea = entity.NewGameArea("surface", 0, 0, constants.SurfaceSize, constants.SurfaceSize)
		e.GameAreas().Store(surfaceArea.Id(), surfaceArea)
		townArea := entity.NewGameArea("town", 0, 0, 50.0, 50.0)
		e.GameAreas().Store(townArea.Id(), townArea)
	})
}

func GetEngine() *game.Engine {
	return e
}

func GetSurfaceArea() *entity.GameArea {
	return surfaceArea
}
