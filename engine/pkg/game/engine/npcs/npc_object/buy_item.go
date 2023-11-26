package npc_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

const (
	TradeDistance = 1.0
)

func (npcObj *NpcObject) BuyItem(e entity.IEngine, charGameObj entity.IGameObject, itemKey string, amount float64) bool {
	playerId := charGameObj.Properties()["player_id"].(int)
	if player, ok := e.Players()[playerId]; ok {
		slots := charGameObj.Properties()["slots"].(map[string]interface{})

		if npcObj.GetDistance(charGameObj) > TradeDistance {
			e.SendSystemMessage("You need to be closer.", player)
			return false
		}

		// get required resources amounts
		resourceKey := npcObj.Properties()["sells"].(map[string]interface{})[itemKey].(map[string]interface{})["resource"].(string)
		resourceAmount := npcObj.Properties()["sells"].(map[string]interface{})[itemKey].(map[string]interface{})["price"].(float64) * amount

		container := e.GameObjects()[slots["back"].(string)]
		// check container has items
		if !container.(entity.IContainerObject).HasItemsKinds(e, map[string]interface{}{
			(resourceKey): resourceAmount,
		}) {
			e.SendSystemMessage("You don't have required resources.", player)
			return false
		}

		// substract resources/money
		if !container.(entity.IContainerObject).RemoveItemsKinds(e, player, map[string]interface{}{
			(resourceKey): resourceAmount,
		}) {
			e.SendSystemMessage("Can't remove required resources.", player)
			return false
		}

		// Create items
		itemObj := e.CreateGameObject(itemKey, charGameObj.X(), charGameObj.Y(), 0.0, charGameObj.Floor(), map[string]interface{}{
			"visible": false,
		})

		// check character has container
		for i := 0; i < int(amount); i++ {
			putInContainer := false
			if (slots["back"] != nil) {
				// put log to container
				putInContainer = container.(entity.IContainerObject).Put(e, player, itemObj.Id(), -1)
			}

			// OR drop item on the ground
			if !putInContainer {
				itemObj.Properties()["visible"] = true
				e.Floors()[itemObj.Floor()].Insert(itemObj)
				storage.GetClient().Updates <- itemObj.Clone()
				e.SendResponseToVisionAreas(e.GameObjects()[player.CharacterGameObjectId], "add_object", map[string]interface{}{
					"object": itemObj,
				})
			}
		}
	} else {
		return false
	}

	return true
}
