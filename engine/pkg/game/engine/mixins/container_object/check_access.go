package container_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (cont *ContainerObject) CheckAccess(e entity.IEngine, player *entity.Player) bool {
	container := cont.gameObj
	return player.CharacterGameObjectId  == container.Properties()["owner_id"]
}
