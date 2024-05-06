package npc_object

import (
	"errors"
	"strings"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (npcObj *NpcObject) SellItem(e entity.IEngine, charGameObj entity.IGameObject, itemKey string, amount float64) (bool, error) {
	playerId := charGameObj.GetProperty("player_id").(int)
	if player, ok := e.Players().Load(playerId); ok {
		slots := charGameObj.GetProperty("slots").(map[string]interface{})
		actions := npcObj.GetProperty("actions")

		if actions == nil || actions.(map[string]interface{})["trade"] == nil {
			e.SendSystemMessage("You can't trade with this NPC", player)
			return false, errors.New("NPC cannot trade")
		}

		if slots["back"] == nil {
			e.SendSystemMessage("You don't have container", player)
			return false, errors.New("Player does not have container")
		}

		if npcObj.GetDistance(charGameObj) > TradeDistance {
			e.SendSystemMessage("You need to be closer.", player)
			return false, errors.New("Player need to be closer to NPC")
		}

		var (
			container entity.IGameObject
			contOk bool
		)
		if container, contOk = e.GameObjects().Load(slots["back"].(string)); !contOk {
			return false, errors.New("Player does not have container")
		}
		itemKind := strings.Split(itemKey, "/")[1]

		// check container has items
		if !container.(entity.IContainerObject).HasItemsKinds(e, map[string]interface{}{
			itemKind: amount,
		}) {
			e.SendSystemMessage("You don't have required resources.", player)
			return false, errors.New("Player does not have required resources")
		}

		// substract resources/money
		if !container.(entity.IContainerObject).RemoveItemsKinds(e, player, map[string]interface{}{
			itemKind: amount,
		}) {
			e.SendSystemMessage("Can't remove required resources.", player)
			return false, errors.New("Player can not remove required resources")
		}

		resourceAmount := npcObj.GetProperty("buys").(map[string]interface{})[itemKey].(map[string]interface{})["price"].(float64) * amount
		resourceKind := npcObj.GetProperty("buys").(map[string]interface{})[itemKey].(map[string]interface{})["resource"].(string)
		resourceKey := "resource/" + resourceKind

		resourceObj := e.CreateGameObject(resourceKey, charGameObj.X(), charGameObj.Y(), 0.0, charGameObj.GameAreaId(), map[string]interface{}{
			"visible": false,
		})

		if isStackable := resourceObj.GetProperty("stackable"); isStackable != nil {
			if isStackable.(bool) {
				resourceObj.SetProperty("amount", resourceAmount)
				return container.(entity.IContainerObject).PutOrDrop(e, charGameObj, resourceObj.Id(), -1), nil
			}
		}

		for i := 0; i < int(resourceAmount); i++ {
			container.(entity.IContainerObject).PutOrDrop(e, charGameObj, resourceObj.Id(), -1)

			// eliminate creating redundant object
			if i == int(resourceAmount)-1 {
				break
			}

			resourceObj = e.CreateGameObject(resourceKey, charGameObj.X(), charGameObj.Y(), 0.0, charGameObj.GameAreaId(), map[string]interface{}{
				"visible": false,
			})
		}

		return true, nil
	}

	return false, errors.New("Player does not exist")
}
