package container_object

import (
	"slices"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

func (cont *ContainerObject) Remove(e entity.IEngine, player *entity.Player, itemId string) bool {
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

	itemPosition := slices.IndexFunc(container.GetProperty("items_ids").([]interface{}), func(id interface{}) bool { return id != nil && id.(string) == itemId })
	if (itemPosition == -1) {
		e.SendSystemMessage("Item is not found in container", player)
	}

	contItemsIds := container.GetProperty("items_ids").([]interface{})
	contItemsIds[itemPosition] = nil
	container.SetProperty("items_ids", contItemsIds)
	container.SetProperty("free_capacity", container.GetProperty("free_capacity").(float64) + 1)
	item.SetProperty("container_id", nil)
	if item.Type() == "container" {
		item.SetProperty("owner_id", nil)
		item.SetProperty("parent_container_id", nil)
	}

	// Save game objects updates to storage
	storage.GetClient().Updates <- container.Clone()
	storage.GetClient().Updates <- item.Clone()

	// Send updates to players
	if charGameObj, charOk := e.GameObjects().Load(player.CharacterGameObjectId); charOk {
		e.SendResponseToVisionAreas(charGameObj, "remove_item_from_container", map[string]interface{}{
			"item_id": itemId,
			"container_id": container.Id(),
			"position": itemPosition,
		})
	}

	return true
}
