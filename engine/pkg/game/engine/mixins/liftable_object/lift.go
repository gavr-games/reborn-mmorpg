package liftable_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/claims"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

// The idea is that character can lift up some items like chests, carry them  and put in another place.
func (obj *LiftableObject) Lift(e entity.IEngine, charGameObj entity.IGameObject) bool {
	playerId := charGameObj.Properties()["player_id"].(int)
	if player, ok := e.Players().Load(playerId); ok {
		item := obj.gameObj

		if item == nil {
			e.SendSystemMessage("Wrong item.", player)
			return false
		}

		if !item.Properties()["liftable"].(bool) {
			e.SendSystemMessage("Wrong item.", player)
			return false
		}
		
		// Check already lifted something
		liftedObjectId, propExists := charGameObj.Properties()["lifted_object_id"]
		if propExists && liftedObjectId != nil {
			e.SendSystemMessage("You are already carrying something.", player)
			return false
		}

		// Check already lifted by someone
		liftedBy, propExists2 := item.Properties()["lifted_by"]
		if propExists2 && liftedBy != nil {
			e.SendSystemMessage("Item is already lifted.", player)
			return false
		}

		// Check is close
		if !item.IsCloseTo(charGameObj) {
			e.SendSystemMessage("You need to be closer.", player)
			return false
		}

		// Check has clasim access
		if !claims.CheckAccess(e, charGameObj, item) {
			e.SendSystemMessage("You don't have an access to this claim.", player)
			return false
		}

		// Update lifted_by and lifted_object_id
		charGameObj.Properties()["lifted_object_id"] = item.Id()
		item.Properties()["lifted_by"] = charGameObj.Id()
		item.Properties()["collidable"] = false
		e.Floors()[item.Floor()].FilteredRemove(item, func(b utils.IBounds) bool {
			return item.Id() == b.(entity.IGameObject).Id()
		})
		item.SetX(charGameObj.X())
		item.SetY(charGameObj.Y())
		e.Floors()[item.Floor()].Insert(item)

		e.SendGameObjectUpdate(charGameObj, "update_object")
		e.SendGameObjectUpdate(item, "update_object")

		e.SendResponseToVisionAreas(charGameObj, "pickup_object", map[string]interface{}{
			"character_id": charGameObj.Id(),
			"id":           item.Id(),
		})
	} else {
		return false
	}

	return true
}
