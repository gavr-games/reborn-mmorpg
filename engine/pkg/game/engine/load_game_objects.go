package engine

import (
	"math/rand"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

func LoadGameObjects(e entity.IEngine, floorSize float64) {
	loadedObjectsCount := storage.GetClient().ReadAllGameObjects(func(gameObj *entity.GameObject) {
		e.GameObjects()[gameObj.Id] = gameObj
		if (gameObj.Floor != -1) {
			e.Floors()[gameObj.Floor].Insert(gameObj)
		}
		if gameObj.Type == "player" {
			playerId := int(gameObj.Properties["player_id"].(float64))
			gameObj.Properties["player_id"] = playerId
			e.Players()[playerId] = &entity.Player{
				Id: playerId,
				CharacterGameObjectId: gameObj.Id,
				VisionAreaGameObjectId: "",
				Client: nil,
				VisibleObjects: make(map[string]bool),
			}
		}
		if gameObj.Type == "mob" {
			// starts go routine, which controls mob
			e.Mobs()[gameObj.Id] = entity.NewMob(e, gameObj.Id)
		}
	})

	// init dump world if no game objects in storage
	if (loadedObjectsCount == 0) {
		// grass
		for x := 0; x < 100; x++ {
			for y := 0; y < 100; y++ {
				CreateGameObject(e, "surface/grass", float64(x), float64(y), 0, nil)
			}
		}
		// rocks
		for i := 0; i < 20; i++ {
			x := 1.0 + rand.Float64() * (99.0 - 1.0)
			y := 1.0 + rand.Float64() * (99.0 - 1.0)
			CreateGameObject(e, "rock/rock_moss", x, y, 0, nil)
		}
		// trees
		for i := 0; i < 20; i++ {
			x := 1.0 + rand.Float64() * (99.0 - 1.0)
			y := 1.0 + rand.Float64() * (99.0 - 1.0)
			CreateGameObject(e, "tree", x, y, 0, nil)
		}
		// eggs
		for i := 0; i < 20; i++ {
			x := 1.0 + rand.Float64() * (99.0 - 1.0)
			y := 1.0 + rand.Float64() * (99.0 - 1.0)
			CreateGameObject(e, "resource/fire_dragon_egg", x, y, 0, map[string]interface{}{
				"visible": true,
			})
		}
	}
}