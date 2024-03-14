package engine

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// Process when player disconnects from the game
func UnregisterClient(e entity.IEngine, client entity.IClient) {
	if player, ok := e.Players().Load(client.GetCharacter().Id); ok {
		if visionAreaGameObj, areaOk := e.GameObjects().Load(player.VisionAreaGameObjectId); areaOk && player.Client != nil {
			//TODO: handle issue this gives panic when closing closed channel
			//close(client.GetSendChannel())
			player.Client = nil
			e.Floors()[visionAreaGameObj.Floor()].FilteredRemove(visionAreaGameObj, func(b utils.IBounds) bool {
				return player.VisionAreaGameObjectId == b.(entity.IGameObject).Id()
			})
			e.GameObjects().Delete(player.VisionAreaGameObjectId)
			if charObj, charOk := e.GameObjects().Load(player.CharacterGameObjectId); charOk {
				charObj.SetProperty("visible", false)
				charObj.(entity.IMovingObject).Stop(e)
				storage.GetClient().Updates <- charObj.Clone()
				e.SendResponseToVisionAreas(charObj, "remove_object", map[string]interface{}{
					"object": charObj,
				})
				// Hide lifted object
				if liftedObjectId := charObj.GetProperty("lifted_object_id"); ok && liftedObjectId != nil {
					if liftedObj, liftedOk := e.GameObjects().Load(liftedObjectId.(string)); liftedOk {
						if liftedObj != nil {
							liftedObj.SetProperty("visible", false)
							storage.GetClient().Updates <- liftedObj.Clone()
							e.SendResponseToVisionAreas(liftedObj, "remove_object", map[string]interface{}{
								"object": liftedObj,
							})
						}
					}
				}
			}
		}
	}
}
