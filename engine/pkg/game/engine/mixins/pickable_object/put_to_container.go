package pickable_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (obj *PickableObject) PutToContainer(e entity.IEngine, containerId string, pos int, player *entity.Player) bool {
	item := obj.gameObj

	//check in container
	currentContainerId := item.GetProperty("container_id")
	if currentContainerId == nil {
		e.SendSystemMessage("Item must be in container.", player)
		return false
	}

	// check containers belongs to character
	if currentContainerId != nil {
		if container, contOk := e.GameObjects().Load(currentContainerId.(string)); contOk {
			if !container.(entity.IContainerObject).CheckAccess(e, player) {
				e.SendSystemMessage("You don't have access to this container", player)
				return false
			}
		} else {
			return false
		}
	}
	if containerTo, containerToOk := e.GameObjects().Load(containerId); containerToOk {
		if !containerTo.(entity.IContainerObject).CheckAccess(e, player) {
			e.SendSystemMessage("You don't have access to this container", player)
			return false
		}
	} else {
		return false
	}

	// remove from container if in container
	if currentContainerId != nil {
		if container, contOk := e.GameObjects().Load(currentContainerId.(string)); contOk {
			if !container.(entity.IContainerObject).Remove(e, player, item.Id()) {
				e.SendSystemMessage("Cannot remove item from container", player)
				return false
			}
		} else {
			return false
		}
	}
	
	// put to container
	if item.GetProperty("container_id") == nil {
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
