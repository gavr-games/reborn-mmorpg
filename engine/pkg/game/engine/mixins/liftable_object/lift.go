package liftable_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/claims"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

// The idea is that character can lift up some items like chests, carry them  and put in another place.
func (obj *LiftableObject) Lift(e entity.IEngine, charGameObj entity.IGameObject) bool {
	playerId := charGameObj.GetProperty("player_id").(int)
	if player, ok := e.Players().Load(playerId); ok {
		item := obj.gameObj

		if item == nil {
			e.SendSystemMessage("Wrong item.", player)
			return false
		}

		if liftable := item.GetProperty("liftable"); liftable == nil || !liftable.(bool) {
			e.SendSystemMessage("Wrong item.", player)
			return false
		}
		
		// Check already lifted something
		
		if liftedObjectId := charGameObj.GetProperty("lifted_object_id"); liftedObjectId != nil {
			e.SendSystemMessage("You are already carrying something.", player)
			return false
		}

		// Check already lifted by someone
		if liftedBy := item.GetProperty("lifted_by"); liftedBy != nil {
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
		itemId := item.Id()
		charGameObj.SetProperty("lifted_object_id", itemId)
		item.SetProperty("lifted_by", charGameObj.Id())
		item.SetProperty("collidable", false)
		gameArea, gaOk := e.GameAreas().Load(item.GameAreaId())
		if gaOk {
			gameArea.FilteredRemove(item, func(b utils.IBounds) bool {
				return itemId == b.(entity.IGameObject).Id()
			})
		}
		item.SetX(charGameObj.X())
		item.SetY(charGameObj.Y())
		if gaOk {
			gameArea.Insert(item)
		}

		e.SendGameObjectUpdate(charGameObj, "update_object")
		e.SendGameObjectUpdate(item, "update_object")

		e.SendResponseToVisionAreas(charGameObj, "pickup_object", map[string]interface{}{
			"character_id": charGameObj.Id(),
			"id":           itemId,
		})
	} else {
		return false
	}

	return true
}
