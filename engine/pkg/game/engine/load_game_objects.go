package engine

import (
	"sync"

	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

func LoadGameObjects(e entity.IEngine) {
	var wg sync.WaitGroup

	loadedObjectsCount := storage.GetClient().ReadAllGameObjects(func(gameObj entity.IGameObject) {
		wg.Add(1)
		go func (gameObj entity.IGameObject) {
			defer wg.Done()
			if (gameObj.Floor() >= 0) {
				e.Floors()[gameObj.Floor()].Insert(gameObj)
			}
			// init player
			if gameObj.Kind() == "player" {
				playerId := int(gameObj.GetProperty("player_id").(float64))
				gameObj.SetProperty("player_id", playerId)
				e.Players().Store(playerId, &entity.Player{
					Id: playerId,
					CharacterGameObjectId: gameObj.Id(),
					VisionAreaGameObjectId: "",
					Client: nil,
				})
				gameObj.SetProperty("visible", false)
			}
			// init effects
			for effectId, effect := range gameObj.Effects() {
				effectMap := utils.CopyMap(effect.(map[string]interface{}))
				effectMap["id"] = effectId
				effectMap["target_id"] = gameObj.Id()
				e.Effects().Store(effectId, effectMap)
			}
			gameObjStruct := e.CreateGameObjectStruct(gameObj)
			e.GameObjects().Store(gameObj.Id(), gameObjStruct)
			// init mob
			if gameObj.Type() == "mob" {
				e.Mobs().Store(gameObj.Id(), gameObjStruct.(entity.IMobObject))
			}
		}(gameObj)
	})

	wg.Wait()

	// init dump world if no game objects in storage
	if (loadedObjectsCount == 0) {
		GenerateWorld(e)
	}
}