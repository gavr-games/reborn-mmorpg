package container_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/claims"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// Containers inside other containers have parent_container_id - in this case we check access to parent container
// All liftable containers stand in real world (chests), not inside character inventory (bags) - for them we need to check claim access
// All containers, that can be put inside character inventory have owner_id - so we need to check if character is owner
func (cont *ContainerObject) CheckAccess(e entity.IEngine, player *entity.Player) bool {
	var (
		charGameObj, parentContainer entity.IGameObject
		charOk, parentContOk bool
	)
	container := cont.gameObj
	parentContainerId := container.GetProperty("parent_container_id")
	if charGameObj, charOk = e.GameObjects().Load(player.CharacterGameObjectId); !charOk {
		return false
	}

	if parentContainerId != nil {
		if parentContainer, parentContOk = e.GameObjects().Load(parentContainerId.(string)); !parentContOk {
			return false
		}
		return parentContainer.(entity.IContainerObject).CheckAccess(e, player)
	} else {
		if liftable := container.GetProperty("liftable"); liftable != nil {
			if liftedBy := container.GetProperty("lifted_by"); liftedBy != nil {
				return liftedBy.(string) == charGameObj.Id()
			} else if liftable.(bool) {
				return container.IsCloseTo(charGameObj) && claims.CheckAccess(e, charGameObj, container)
			}
		}
		return player.CharacterGameObjectId  == container.GetProperty("owner_id")
	}
}
