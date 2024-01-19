package container_object

import (
	"slices"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/serializers"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

// position: -1 for any empty slot
func (cont *ContainerObject) Put(e entity.IEngine, player *entity.Player, itemId string, position int) bool {
	container := cont.gameObj
	item := e.GameObjects()[itemId]

	if !cont.CheckAccess(e, player) {
		e.SendSystemMessage("You don't have access to this container", player)
		return false
	}

	if item.Type() == "container" && item.Properties()["max_capacity"].(float64) >= container.Properties()["max_capacity"].(float64) {
		e.SendSystemMessage("Container is too big to put it here.", player)
		return false
	}

	itemStackable := false
	if value, ok := item.Properties()["stackable"]; ok {
		itemStackable = value.(bool)
	}

	freePosition := position
	if position == -1 {
		if itemStackable {
			existingItem := cont.GetItemKind(e, item.Kind())
			if existingItem != nil {
				existingItem.Properties()["amount"] = existingItem.Properties()["amount"].(float64) + item.Properties()["amount"].(float64)
				e.SendGameObjectUpdate(existingItem, "update_object")
				return true
			}
		}

		if container.Properties()["free_capacity"] == 0.0 {
			// Also search free space inside sub-containers
			// TODO: mute messages for sub containers
			subItemIds := container.Properties()["items_ids"].([]interface{})
			for _, subItemId := range subItemIds {
				if subItemId != nil {
					subItem := e.GameObjects()[subItemId.(string)]
					if subItem.Type() == "container" {
						if subItem.(entity.IContainerObject).Put(e, player, itemId, position) {
							return true
						}
					}
				}
			}
			e.SendSystemMessage("This container is full.", player)
			return false
		}

		freePosition = slices.IndexFunc(container.Properties()["items_ids"].([]interface{}), func(id interface{}) bool { return id == nil })
	} else {
		if container.Properties()["items_ids"].([]interface{})[position] == nil {
			freePosition = position
		} else {
			e.SendSystemMessage("This slot inside the container is already occupied.", player)
			return false
		}
	}

	// Modify game objects
	container.Properties()["items_ids"].([]interface{})[freePosition] = itemId
	container.Properties()["free_capacity"] = container.Properties()["free_capacity"].(float64) - 1.0
	item.Properties()["container_id"] = container.Id()
	item.Properties()["visible"] = false
	if item.Type() == "container" {
		item.Properties()["owner_id"] = container.Properties()["owner_id"]
		item.Properties()["parent_container_id"] = container.Id()
	}

	// Save game objects updates to storage
	storage.GetClient().Updates <- container.Clone()
	storage.GetClient().Updates <- item.Clone()

	// Send updates to players
	e.SendResponseToVisionAreas(e.GameObjects()[player.CharacterGameObjectId], "put_item_to_container", map[string]interface{}{
		"item":         serializers.GetInfo(e.GameObjects(), item),
		"container_id": container.Id(),
		"position":     freePosition,
	})

	return true
}
