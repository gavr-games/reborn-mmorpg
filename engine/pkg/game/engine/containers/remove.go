package containers

import (
	"slices"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

func Remove(e entity.IEngine, player *entity.Player, containerId string, itemId string) bool {
	container := e.GameObjects()[containerId]
	item := e.GameObjects()[itemId]

	if !CheckAccess(e, player, container) {
		e.SendSystemMessage("You don't have access to this container", player)
		return false
	}

	itemPosition := slices.IndexFunc(container.Properties["items_ids"].([]interface{}), func(id interface{}) bool { return id != nil && id.(string) == itemId })
	if (itemPosition == -1) {
		e.SendSystemMessage("Item is not found in container", player)
	}

	container.Properties["items_ids"].([]interface{})[itemPosition] = nil
	container.Properties["free_capacity"] = container.Properties["free_capacity"].(float64) + 1
	item.Properties["container_id"] = nil

	// Save game objects updates to storage
	storage.GetClient().Updates <- container
	storage.GetClient().Updates <- item

	// Send updates to players
	e.SendResponseToVisionAreas(e.GameObjects()[player.CharacterGameObjectId], "remove_item_from_container", map[string]interface{}{
		"item_id": itemId,
		"container_id": containerId,
		"position": itemPosition,
	})

	return true
}
