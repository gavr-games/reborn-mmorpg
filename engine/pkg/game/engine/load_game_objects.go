package engine

import (
	"math/rand"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

func LoadGameObjects(e IEngine, floorSize float64) {
	loadedObjectsCount := storage.GetClient().ReadAllGameObjects(func(gameObj *entity.GameObject) {
		e.GameObjects()[gameObj.Id] = gameObj
		e.Floors()[gameObj.Floor].Insert(gameObj)
		if gameObj.Type == "player" {
			playerId := gameObj.Properties["player_id"].(float64)
			e.Players()[int(playerId)] = &entity.Player{
				Id: int(playerId),
				CharacterGameObjectId: gameObj.Id,
				VisionAreaGameObjectId: "",
				Client: nil,
				VisibleObjects: make(map[string]bool),
			}
		}
	})

	// init dump world if no game objects in storage
	if (loadedObjectsCount == 0) {
		// grass
		for x := 0; x < 100; x++ {
			for y := 0; y < 100; y++ {
				// + 0.5 because we want to place the center point
				gameObj := CreateGameObject("surface/grass", float64(x) + 0.5, float64(y) + 0.5, nil)
				gameObj.Floor = 0
				e.GameObjects()[gameObj.Id] = gameObj
				e.Floors()[gameObj.Floor].Insert(gameObj)
			}
		}
		// rocks
		for i := 0; i < 20; i++ {
			x := 1.0 + rand.Float64() * (99.0 - 1.0)
			y := 1.0 + rand.Float64() * (99.0 - 1.0)
			gameObj := CreateGameObject("rock/rock_moss", x, y, nil)
			gameObj.Floor = 0
			e.GameObjects()[gameObj.Id] = gameObj
			e.Floors()[gameObj.Floor].Insert(gameObj)
		}
		// trees
		for i := 0; i < 20; i++ {
			x := 1.0 + rand.Float64() * (99.0 - 1.0)
			y := 1.0 + rand.Float64() * (99.0 - 1.0)
			gameObj := CreateGameObject("tree", x, y, nil)
			gameObj.Floor = 0
			e.GameObjects()[gameObj.Id] = gameObj
			e.Floors()[gameObj.Floor].Insert(gameObj)
		}
	}
}