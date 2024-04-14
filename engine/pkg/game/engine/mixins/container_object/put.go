package container_object

import (
	"slices"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/serializers"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

// position: -1 for any empty slot
func (cont *ContainerObject) Put(e entity.IEngine, player *entity.Player, itemId string, position int) bool {
	var (
		item entity.IGameObject
		itemOk bool
	)
	container := cont.gameObj
	if item, itemOk = e.GameObjects().Load(itemId); !itemOk {
		return false
	}

	if !cont.CheckAccess(e, player) {
		e.SendSystemMessage("You don't have access to this container", player)
		return false
	}

	if item.Type() == "container" && item.GetProperty("max_capacity").(float64) >= container.GetProperty("max_capacity").(float64) {
		e.SendSystemMessage("Container is too big to put it here.", player)
		return false
	}

	itemStackable := false
	if stackable := item.GetProperty("stackable"); stackable != nil {
		itemStackable = stackable.(bool)
	}

	freePosition := position
	if position == -1 {
		if itemStackable {
			existingItem := cont.GetItemKind(e, item.Kind())
			if existingItem != nil {
				existingItem.SetProperty("amount", existingItem.GetProperty("amount").(float64) + item.GetProperty("amount").(float64))
				e.SendGameObjectUpdate(existingItem, "update_object")
				e.RemoveGameObject(item)
				return true
			}
		}

		if container.GetProperty("free_capacity") == 0.0 {
			// Also search free space inside sub-containers
			// TODO: mute messages for sub containers
			subItemIds := container.GetProperty("items_ids").([]interface{})
			for _, subItemId := range subItemIds {
				if subItemId != nil {
					if subItem, subItemOk := e.GameObjects().Load(subItemId.(string)); subItemOk {
						if subItem.Type() == "container" {
							if subItem.(entity.IContainerObject).Put(e, player, itemId, position) {
								return true
							}
						}
					}
				}
			}
			e.SendSystemMessage("This container is full.", player)
			return false
		}

		freePosition = slices.IndexFunc(container.GetProperty("items_ids").([]interface{}), func(id interface{}) bool { return id == nil })
	} else {
		if container.GetProperty("items_ids").([]interface{})[position] == nil {
			freePosition = position
		} else {
			e.SendSystemMessage("This slot inside the container is already occupied.", player)
			return false
		}
	}

	// Modify game objects
	contItemsIds := container.GetProperty("items_ids").([]interface{})
	contItemsIds[freePosition] = itemId
	container.SetProperty("items_ids", contItemsIds)
	container.SetProperty("free_capacity", container.GetProperty("free_capacity").(float64) - 1.0)
	item.SetProperty("container_id", container.Id())
	item.SetProperty("visible", false)
	if item.Type() == "container" {
		item.SetProperty("owner_id", container.GetProperty("owner_id"))
		item.SetProperty("parent_container_id", container.Id())
	}

	// Save game objects updates to storage
	storage.GetClient().Updates <- container.Clone()
	storage.GetClient().Updates <- item.Clone()

	// Send updates to players
	if charGameObj, charOk := e.GameObjects().Load(player.CharacterGameObjectId); charOk {
		e.SendResponseToVisionAreas(charGameObj, "put_item_to_container", map[string]interface{}{
			"item":         serializers.GetInfo(e, item),
			"container_id": container.Id(),
			"position":     freePosition,
		})
	}

	return true
}
