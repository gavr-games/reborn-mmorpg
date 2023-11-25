package items

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func PutToContainer(e entity.IEngine, containerId string, pos int, itemId string, player *entity.Player) bool {
	item := e.GameObjects()[itemId]

	if item == nil {
		e.SendSystemMessage("Wrong item.", player)
		return false
	}

	//check in container
	if (item.Properties()["container_id"] == nil) {
		e.SendSystemMessage("Item must be in container.", player)
		return false
	}

	// check container belongs to character
	if (item.Properties()["container_id"] != nil) {
		container := e.GameObjects()[item.Properties()["container_id"].(string)]
		if !container.(entity.IContainerObject).CheckAccess(e, player) {
			e.SendSystemMessage("You don't have access to this container", player)
			return false
		}
	}

	// remove from container if in container
	if (item.Properties()["container_id"] != nil) {
		container := e.GameObjects()[item.Properties()["container_id"].(string)]
		if !container.(entity.IContainerObject).Remove(e, player, itemId) {
			e.SendSystemMessage("Cannot remove item from container", player)
			return false
		}
	}
	
	// put to container
	if (item.Properties()["container_id"] == nil) {
		containerTo := e.GameObjects()[containerId]
		if !containerTo.(entity.IContainerObject).Put(e, player, itemId, pos) {
			e.SendSystemMessage("Cannot put item to container", player)
			return false
		}
	}

	return true
}
