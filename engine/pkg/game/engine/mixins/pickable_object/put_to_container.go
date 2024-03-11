package pickable_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (obj *PickableObject) PutToContainer(e entity.IEngine, containerId string, pos int, player *entity.Player) bool {
	item := obj.gameObj

	//check in container
	if (item.Properties()["container_id"] == nil) {
		e.SendSystemMessage("Item must be in container.", player)
		return false
	}

	// check container belongs to character
	if (item.Properties()["container_id"] != nil) {
		if container, contOk := e.GameObjects().Load(item.Properties()["container_id"].(string)); contOk {
			if !container.(entity.IContainerObject).CheckAccess(e, player) {
				e.SendSystemMessage("You don't have access to this container", player)
				return false
			}
		} else {
			return false
		}
	}

	// remove from container if in container
	if (item.Properties()["container_id"] != nil) {
		if container, contOk := e.GameObjects().Load(item.Properties()["container_id"].(string)); contOk {
			if !container.(entity.IContainerObject).Remove(e, player, item.Id()) {
				e.SendSystemMessage("Cannot remove item from container", player)
				return false
			}
		} else {
			return false
		}
	}
	
	// put to container
	if (item.Properties()["container_id"] == nil) {
		if containerTo, containerToOk := e.GameObjects().Load(containerId); containerToOk {
			if !containerTo.(entity.IContainerObject).Put(e, player, item.Id(), pos) {
				e.SendSystemMessage("Cannot put item to container", player)
				return false
			}
		} else {
			return false
		}
	}

	return true
}
