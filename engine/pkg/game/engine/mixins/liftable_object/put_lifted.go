package liftable_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

const (
	PutDistance = 0.5
)

func (obj *LiftableObject) PutLifted(e entity.IEngine, charGameObj entity.IGameObject, x float64, y float64, rotation float64) bool {
	playerId := charGameObj.Properties()["player_id"].(int)
	if player, ok := e.Players().Load(playerId); ok {
		item := obj.gameObj

		if item == nil {
			e.SendSystemMessage("Wrong item.", player)
			return false
		}

		if charGameObj.Properties()["lifted_object_id"].(string) != item.Id() {
			e.SendSystemMessage("Wrong item.", player)
			return false
		}

		// Create clone
		clone := item.Clone()
		clone.SetX(x)
		clone.SetY(y)
		clone.Rotate(rotation)

		// Check distance
		if charGameObj.GetDistance(clone) > PutDistance {
			e.SendSystemMessage("You need to be closer.", player)
			return false
		}

		// Check not intersecting with another objects
		possibleCollidableObjects := e.Floors()[charGameObj.Floor()].RetrieveIntersections(utils.Bounds{
			X:      x,
			Y:      y,
			Width:  clone.Width(),
			Height: clone.Height(),
		})

		if len(possibleCollidableObjects) > 0 {
			for _, val := range possibleCollidableObjects {
				gameObj := val.(entity.IGameObject)
				if collidable, ok := gameObj.Properties()["collidable"]; ok {
					if collidable.(bool) {
						e.SendSystemMessage("Cannot put it here. There is something in the way.", player)
						return false
					}
				}
			}
		}

		// Update Objects
		charGameObj.Properties()["lifted_object_id"] = nil
		item.Properties()["lifted_by"] = nil
		item.Properties()["collidable"] = true
		e.Floors()[item.Floor()].FilteredRemove(item, func(b utils.IBounds) bool {
			return item.Id() == b.(entity.IGameObject).Id()
		})
		item.SetX(x)
		item.SetY(y)
		item.Rotate(rotation)
		e.Floors()[item.Floor()].Insert(item)

		e.SendGameObjectUpdate(charGameObj, "update_object")
		e.SendGameObjectUpdate(item, "update_object")
	} else {
		return false
	}

	return true
}