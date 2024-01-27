package character_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/players"
)

func (charGameObj *CharacterObject) Move(e entity.IEngine, newX float64, newY float64) {
	playerId := charGameObj.Properties()["player_id"].(int)
	if player, ok := e.Players()[playerId]; ok {
		visionAreaGameObj := e.GameObjects()[player.VisionAreaGameObjectId]

		e.Floors()[charGameObj.Floor()].FilteredRemove(charGameObj, func(b utils.IBounds) bool {
			return charGameObj.Id() == b.(entity.IGameObject).Id()
		})
		dx := newX - charGameObj.X()
		dy := newY - charGameObj.Y()
		charGameObj.SetX(newX)
		charGameObj.SetY(newY)
		e.Floors()[charGameObj.Floor()].Insert(charGameObj)

		// Update vision area game object
		e.Floors()[visionAreaGameObj.Floor()].FilteredRemove(visionAreaGameObj, func(b utils.IBounds) bool {
			return visionAreaGameObj.Id() == b.(entity.IGameObject).Id()
		})
		visionAreaGameObj.SetX(visionAreaGameObj.X() + dx)
		visionAreaGameObj.SetY(visionAreaGameObj.Y() + dy)
		e.Floors()[visionAreaGameObj.Floor()].Insert(visionAreaGameObj)

		// Update lifted item
		liftedObjectId, propExists := charGameObj.Properties()["lifted_object_id"]
		if propExists && liftedObjectId != nil {
			liftedObj := e.GameObjects()[liftedObjectId.(string)]
			if liftedObj != nil {
				e.Floors()[liftedObj.Floor()].FilteredRemove(liftedObj, func(b utils.IBounds) bool {
					return liftedObj.Id() == b.(entity.IGameObject).Id()
				})
				liftedObj.SetX(charGameObj.X())
				liftedObj.SetY(charGameObj.Y())
				e.Floors()[liftedObj.Floor()].Insert(liftedObj)
			}
		}

		// determine new and old visible objects, send updates to client
		visibleObjects := players.GetVisibleObjects(e, player)
		for id, _ := range player.VisibleObjects {
			player.VisibleObjects[id] = false
		}
		// send add new visible objects
		// TODO: add serializers to minimize traffic
		var addObjects []entity.IGameObject
		for _, val := range visibleObjects {
			if _, ok := player.VisibleObjects[val.(entity.IGameObject).Id()]; !ok {
				addObjects = append(addObjects, val.(entity.IGameObject))
			}
			player.VisibleObjects[val.(entity.IGameObject).Id()] = true
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