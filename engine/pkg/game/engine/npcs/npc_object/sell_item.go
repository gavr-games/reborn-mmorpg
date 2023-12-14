package npc_object

import (
	"strings"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

func (npcObj *NpcObject) SellItem(e entity.IEngine, charGameObj entity.IGameObject, itemKey string, amount float64) bool {
	playerId := charGameObj.Properties()["player_id"].(int)
	if player, ok := e.Players()[playerId]; ok {
		slots := charGameObj.Properties()["slots"].(map[string]interface{})

		if npcObj.GetDistance(charGameObj) > TradeDistance {
			e.SendSystemMessage("You need to be closer.", player)
			return false
		}

		container := e.GameObjects()[slots["back"].(string)]
		itemKind := strings.Split(itemKey, "/")[1]

		// check container has items
		if !container.(entity.IContainerObject).HasItemsKinds(e, map[string]interface{}{
			itemKind: amount,
		}) {
			e.SendSystemMessage("You don't have required resources.", player)
			return false
		}

		// substract resources/money
		if !container.(entity.IContainerObject).RemoveItemsKinds(e, player, map[string]interface{}{
			itemKind: amount,
		}) {
			e.SendSystemMessage("Can't remove required resources.", player)
			return false
		}

		resourceAmount := npcObj.Properties()["buys"].(map[string]interface{})[itemKey].(map[string]interface{})["price"].(float64) * amount
		resourceKind := npcObj.Properties()["buys"].(map[string]interface{})[itemKey].(map[string]interface{})["resource"].(string)
		resourceKey := "resource/" + resourceKind

		resourceObject := e.CreateGameObject(resourceKey, charGameObj.X(), charGameObj.Y(), 0.0, charGameObj.Floor(), map[string]interface{}{
			"visible": false,
		})

		if container.(entity.IContainerObject).AddResource(e, player, resourceObject, resourceAmount) {
			return true
		}

		resourceStackable := false
		if value, ok := resourceObject.Properties()["stackable"]; ok {
			resourceStackable = value.(bool)
		}

		for i := 0; i < int(resourceAmount); i++ {
			resourceObject.Properties()["visible"] = true
			e.Floors()[resourceObject.Floor()].Insert(resourceObject)
			storage.GetClient().Updates <- resourceObject.Clone()
			e.SendResponseToVisionAreas(e.GameObjects()[player.CharacterGameObjectId], "add_object", map[string]interface{}{
				"object": resourceObject,
			})

			if resourceStackable {
				break
			}

			resourceObject = e.CreateGameObject(resourceKey, charGameObj.X(), charGameObj.Y(), 0.0, charGameObj.Floor(), map[string]interface{}{
				"visible": false,
			})
		}

		return true
	}

	return false
}
