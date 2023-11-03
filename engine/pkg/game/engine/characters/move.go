package characters

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/players"
)

func Move(e entity.IEngine, charGameObj *entity.GameObject, newX float64, newY float64) {
	playerId := charGameObj.Properties["player_id"].(int)
	if player, ok := e.Players()[playerId]; ok {
		visionAreaGameObj := e.GameObjects()[player.VisionAreaGameObjectId]

		e.Floors()[charGameObj.Floor].FilteredRemove(charGameObj, func(b utils.IBounds) bool {
			return charGameObj.Id == b.(*entity.GameObject).Id
		})
		dx := newX - charGameObj.X
		dy := newY - charGameObj.Y
		charGameObj.X = newX
		charGameObj.Y = newY
		charGameObj.Properties["x"] = newX
		charGameObj.Properties["y"] = newY
		e.Floors()[charGameObj.Floor].Insert(charGameObj)
		// Update vision area game object
		e.Floors()[visionAreaGameObj.Floor].FilteredRemove(visionAreaGameObj, func(b utils.IBounds) bool {
			return visionAreaGameObj.Id == b.(*entity.GameObject).Id
		})
		visionAreaGameObj.X += dx
		visionAreaGameObj.Y += dy
		visionAreaGameObj.Properties["x"] = visionAreaGameObj.Properties["x"].(float64) + dx
		visionAreaGameObj.Properties["y"] = visionAreaGameObj.Properties["y"].(float64) + dy
		e.Floors()[visionAreaGameObj.Floor].Insert(visionAreaGameObj)

		// determine new and old visible objects, send updates to client
		visibleObjects := players.GetVisibleObjects(e, player)
		for id, _ := range player.VisibleObjects {
			player.VisibleObjects[id] = false
		}
		// send add new visible objects
		// TODO: add serializers to minimize traffic
		var addObjects []*entity.GameObject
		for _, val := range visibleObjects {
			if _, ok := player.VisibleObjects[val.(*entity.GameObject).Id]; !ok {
				addObjects = append(addObjects, val.(*entity.GameObject))
			}
			player.VisibleObjects[val.(*entity.GameObject).Id] = true
		}
		if len(addObjects) > 0 {
			e.SendResponse("add_objects", map[string]interface{}{
				"objects": addObjects,
			}, player)
		}
		// send remove old visible objects
		var removeObjects []map[string]interface{}
		for id, visible := range player.VisibleObjects {
			if !visible {
				removeObjects = append(removeObjects, map[string]interface{}{
					"Id": id,
				})
				delete(player.VisibleObjects, id)
			}
		}
		if len(removeObjects) > 0 {
			e.SendResponse("remove_objects", map[string]interface{}{
				"objects": removeObjects,
			}, player)
		}
	}
}