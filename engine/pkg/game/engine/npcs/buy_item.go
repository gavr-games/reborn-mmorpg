package npcs

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/containers"
)

const (
	TradeDistance = 1.0
)

func BuyItem(e entity.IEngine, charGameObj *entity.GameObject, npcId string, itemKey string, amount float64) bool {
	playerId := charGameObj.Properties["player_id"].(int)
	if player, ok := e.Players()[playerId]; ok {
		npcObj := e.GameObjects()[npcId]
		slots := charGameObj.Properties["slots"].(map[string]interface{})

		if npcObj == nil {
			e.SendSystemMessage("NPC does not exist.", player)
			return false
		}

		if game_objects.GetDistance(npcObj, charGameObj) > TradeDistance {
			e.SendSystemMessage("You need to be closer.", player)
			return false
		}

		// get required resources amounts
		resourceKey := npcObj.Properties["sells"].(map[string]interface{})[itemKey].(map[string]interface{})["resource"].(string)
		resourceAmount := npcObj.Properties["sells"].(map[string]interface{})[itemKey].(map[string]interface{})["price"].(float64) * amount

		// check container has items
		if !containers.HasItemsKinds(e, slots["back"].(string), map[string]interface{}{
			(resourceKey): resourceAmount,
		}) {
			e.SendSystemMessage("You don't have required resources.", player)
			return false
		}

		// substract resources/money
		if !containers.RemoveItemsKinds(e, player, slots["back"].(string), map[string]interface{}{
			(resourceKey): resourceAmount,
		}) {
			e.SendSystemMessage("Can't remove required resources.", player)
			return false
		}

		// Create items
		itemObj := e.CreateGameObject(itemKey, charGameObj.X, charGameObj.Y, 0.0, charGameObj.Floor, map[string]interface{}{
			"visible": false,
		})

		// check character has container
		for i := 0; i < int(amount); i++ {
			putInContainer := false
			if (slots["back"] != nil) {
				// put log to container
				putInContainer = containers.Put(e, player, slots["back"].(string), itemObj.Id, -1)
			}

			// OR drop item on the ground
			if !putInContainer {
				itemObj.Properties["visible"] = true
				e.Floors()[itemObj.Floor].Insert(itemObj)
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
