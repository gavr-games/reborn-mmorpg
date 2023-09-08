package containers

import (
	"slices"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

// position: -1 for any empty slot
func Put(e entity.IEngine, player *entity.Player, containerId string, itemId string, position int) bool {
	container := e.GameObjects()[containerId]
	item := e.GameObjects()[itemId]

	if container.Properties["free_capacity"] == 0 {
		e.SendResponse("add_message", map[string]interface{}{
			"type": "system",
			"message": "This container is full.",
		}, player)
		return false
	}

	if !CheckAccess(e, player, container) {
		e.SendResponse("add_message", map[string]interface{}{
			"type": "system",
			"message": "You don't have access to this container",
		}, player)
		return false
	}

	freePosition := position
	if position == -1 {
		freePosition = slices.IndexFunc(container.Properties["items_ids"].([]string), func(id string) bool { return id == "" })
	} else {
		if (container.Properties["items_ids"].([]string)[position] == "") {
			freePosition = position
		} else {
			e.SendResponse("add_message", map[string]interface{}{
				"type": "system",
				"message": "This slot inside the container is already occupied .",
			}, player)
			return false
		}
	}

	// Modify game objects
	container.Properties["items_ids"].([]string)[freePosition] = itemId
	container.Properties["free_capacity"] = container.Properties["free_capacity"].(int) - 1
	item.Properties["container_id"] = containerId
	item.Properties["visible"] = false

	// Save game objects updates to storage
	storage.GetClient().Updates <- container
	storage.GetClient().Updates <- item

	// Send updates to players
	e.SendResponseToVisionAreas(e.GameObjects()[player.CharacterGameObjectId], "put_item_to_container", map[string]interface{}{
		"item_id": itemId,
		"container_id": containerId,
		"position": freePosition,
	})

	return true
}
