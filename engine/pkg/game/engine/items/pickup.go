package items

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

func Pickup(e entity.IEngine, itemId string, player *entity.Player) bool {
	item := e.GameObjects()[itemId]
	charGameObj := e.GameObjects()[player.CharacterGameObjectId]
	slots := charGameObj.Properties()["slots"].(map[string]interface{})

	if item == nil {
		e.SendSystemMessage("Wrong item.", player)
		return false
	}

	// intersects with character
	itemBounds := utils.Bounds{
		X:      item.X(),
		Y:      item.Y(),
		Width:  item.Width(),
		Height: item.Height(),
	}
	if !charGameObj.Intersects(itemBounds) {
		e.SendSystemMessage("You are too far away.", player)
		return false
	}

	// not in another container
	if (item.Properties()["container_id"] != nil) {
		e.SendSystemMessage("Item is already in another container.", player)
		return false
	}

	// check character has container
	if (slots["back"] == nil) {
		e.SendSystemMessage("You don't have container to put item to.", player)
		return false
	}

	// put to container
	itemStackable := false
	if value, ok := item.Properties()["stackable"]; ok {
		itemStackable = value.(bool)
	}
	if (item.Properties()["container_id"] == nil) {
		container := e.GameObjects()[slots["back"].(string)]
		performPut := true
		if itemStackable { // add amount to existing stackable item
			existingItem := container.(entity.IContainerObject).GetItemKind(e, item.Properties()["kind"].(string))
			if existingItem != nil {
				existingItem.Properties()["amount"] = existingItem.Properties()["amount"].(float64) + item.Properties()["amount"].(float64)
				performPut = false
				storage.GetClient().Updates <- existingItem.Clone()
				e.SendResponse("update_object", map[string]interface{}{
					"object": existingItem.Clone(),
				}, player)
			}
		}
		if performPut {
			if !container.(entity.IContainerObject).Put(e, player, itemId, -1) {
				return false
			}
		}
	}

	// remove from world
	e.Floors()[item.Floor()].FilteredRemove(item, func(b utils.IBounds) bool {
		return itemId == b.(entity.IGameObject).Id()
	})
	item.Properties()["visible"] = false

	storage.GetClient().Updates <- item.Clone()

	e.SendResponseToVisionAreas(e.GameObjects()[player.CharacterGameObjectId], "pickup_object", map[string]interface{}{
		"character_id": charGameObj.Id(),
		"id": itemId,
	})
	e.SendResponseToVisionAreas(e.GameObjects()[player.CharacterGameObjectId], "remove_object", map[string]interface{}{
		"object": map[string]interface{}{
			"Id": itemId,
		},
	})

	return true
}
