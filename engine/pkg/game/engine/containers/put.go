package containers

import (
	"slices"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/serializers"
)

// position: -1 for any empty slot
func Put(e entity.IEngine, player *entity.Player, containerId string, itemId string, position int) bool {
	container := e.GameObjects()[containerId]
	item := e.GameObjects()[itemId]

	if container.Properties["free_capacity"] == 0.0 {
		e.SendSystemMessage("This container is full.", player)
		return false
	}

	if !CheckAccess(e, player, container) {
		e.SendSystemMessage("You don't have access to this container", player)
		return false
	}

	//TODO: also search free space inside sub-containers
	freePosition := position
	if position == -1 {
		freePosition = slices.IndexFunc(container.Properties["items_ids"].([]interface{}), func(id interface{}) bool { return id == nil })
	} else {
		if (container.Properties["items_ids"].([]interface{})[position] == nil) {
			freePosition = position
		} else {
			e.SendSystemMessage("This slot inside the container is already occupied.", player)
			return false
		}
	}

	// Modify game objects
	container.Properties["items_ids"].([]interface{})[freePosition] = itemId
	container.Properties["free_capacity"] = container.Properties["free_capacity"].(float64) - 1.0
	item.Properties["container_id"] = containerId
	item.Properties["visible"] = false

	// Save game objects updates to storage
	storage.GetClient().Updates <- container.Clone()
	storage.GetClient().Updates <- item.Clone()

	// Send updates to players
	e.SendResponseToVisionAreas(e.GameObjects()[player.CharacterGameObjectId], "put_item_to_container", map[string]interface{}{
		"item": serializers.GetInfo(e.GameObjects(), item),
		"container_id": containerId,
		"position": freePosition,
	})

	return true
}
