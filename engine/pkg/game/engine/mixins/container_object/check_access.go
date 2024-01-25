package container_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/claims"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// Containers inside other containers have parent_container_id - in this case we check access to parent container
// All liftable containers stand in real world (chests), not inside character inventory (bags) - for them we need to check claim access
// All containers, that can be put inside character inventory have owner_id - so we need to check if character is owner
func (cont *ContainerObject) CheckAccess(e entity.IEngine, player *entity.Player) bool {
	container := cont.gameObj
	parentContainerId := container.Properties()["parent_container_id"]
	charGameObj := e.GameObjects()[player.CharacterGameObjectId]

	if parentContainerId != nil {
		return e.GameObjects()[parentContainerId.(string)].(entity.IContainerObject).CheckAccess(e, player)
	} else {
		if liftable, ok := container.Properties()["liftable"]; ok {
			if liftable.(bool) {
				return container.IsCloseTo(charGameObj) && claims.CheckAccess(e, charGameObj, container)
			}
		}
		return player.CharacterGameObjectId  == container.Properties()["owner_id"]
	}
}
