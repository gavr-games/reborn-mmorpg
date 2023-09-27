package engine

import (
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
		GenerateWorld(e, floorSize)
	}
}