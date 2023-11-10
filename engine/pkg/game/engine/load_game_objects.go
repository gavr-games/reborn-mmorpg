package engine

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mobs"
)

func LoadGameObjects(e entity.IEngine) {
	loadedObjectsCount := storage.GetClient().ReadAllGameObjects(func(gameObj *entity.GameObject) {
		e.GameObjects()[gameObj.Id] = gameObj
		if (gameObj.Floor >= 0) {
			e.Floors()[gameObj.Floor].Insert(gameObj)
		}
		// init player
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
		// init effects
		for effectId, effect := range gameObj.Effects {
			e.Effects()[effectId] = utils.CopyMap(effect.(map[string]interface{}))
			e.Effects()[effectId]["id"] = effectId
			e.Effects()[effectId]["target_id"] = gameObj.Id
		}
		// init mob
		if gameObj.Type == "mob" {
			e.Mobs()[gameObj.Id] = mobs.NewMob(e, gameObj.Id)
		}
	})

	// init dump world if no game objects in storage
	if (loadedObjectsCount == 0) {
		GenerateWorld(e)
	}
}