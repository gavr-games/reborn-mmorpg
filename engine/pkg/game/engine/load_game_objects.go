package engine

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

func LoadGameObjects(e entity.IEngine) {
	loadedObjectsCount := storage.GetClient().ReadAllGameObjects(func(gameObj entity.IGameObject) {
		if (gameObj.Floor() >= 0) {
			e.Floors()[gameObj.Floor()].Insert(gameObj)
		}
		// init player
		if gameObj.Type() == "player" {
			playerId := int(gameObj.Properties()["player_id"].(float64))
			gameObj.Properties()["player_id"] = playerId
			e.Players()[playerId] = &entity.Player{
				Id: playerId,
				CharacterGameObjectId: gameObj.Id(),
				VisionAreaGameObjectId: "",
				Client: nil,
				VisibleObjects: make(map[string]bool),
			}
		}
		// init effects
		for effectId, effect := range gameObj.Effects() {
			e.Effects()[effectId] = utils.CopyMap(effect.(map[string]interface{}))
			e.Effects()[effectId]["id"] = effectId
			e.Effects()[effectId]["target_id"] = gameObj.Id()
		}
		e.GameObjects()[gameObj.Id()] = e.CreateGameObjectStruct(gameObj)
		// init mob
		if gameObj.Type() == "mob" {
			e.Mobs()[gameObj.Id()] = e.GameObjects()[gameObj.Id()].(entity.IMobObject)
		}
	})

	// init dump world if no game objects in storage
	if (loadedObjectsCount == 0) {
		GenerateWorld(e)
	}
}