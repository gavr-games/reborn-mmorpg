package npc_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (npcObj *NpcObject) BuyItem(e entity.IEngine, charGameObj entity.IGameObject, itemKey string, amount float64) bool {
	playerId := charGameObj.Properties()["player_id"].(int)
	if player, ok := e.Players().Load(playerId); ok {
		slots := charGameObj.Properties()["slots"].(map[string]interface{})

		if slots["back"] == nil {
			e.SendSystemMessage("You don't have container", player)
			return false
		}

		if npcObj.GetDistance(charGameObj) > TradeDistance {
			e.SendSystemMessage("You need to be closer.", player)
			return false
		}

		// get required resources amounts
		resourceKey := npcObj.Properties()["sells"].(map[string]interface{})[itemKey].(map[string]interface{})["resource"].(string)
		resourceAmount := npcObj.Properties()["sells"].(map[string]interface{})[itemKey].(map[string]interface{})["price"].(float64) * amount

		var (
			container entity.IGameObject
			contOk bool
		)
		if container, contOk = e.GameObjects().Load(slots["back"].(string)); !contOk {
			return false
		}
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

		if isStackable, ok := itemObj.Properties()["stackable"]; ok {
			if isStackable.(bool) {
				itemObj.Properties()["amount"] = amount
				return container.(entity.IContainerObject).PutOrDrop(e, charGameObj, itemObj.Id(), -1)
			}
		}

		for i := 0; i < int(amount); i++ {
			container.(entity.IContainerObject).PutOrDrop(e, charGameObj, itemObj.Id(), -1)

			// eliminate creating redundant object
			if i == int(resourceAmount)-1 {
				break
			}

			itemObj = e.CreateGameObject(itemKey, charGameObj.X(), charGameObj.Y(), 0.0, charGameObj.Floor(), map[string]interface{}{
				"visible": false,
			})
		}

		return true
	}

	return false
}
