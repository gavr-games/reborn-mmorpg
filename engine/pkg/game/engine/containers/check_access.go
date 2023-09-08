package containers

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func CheckAccess(e entity.IEngine, player *entity.Player, container *entity.GameObject) bool {
	return player.CharacterGameObjectId  == container.Properties["owner_id"]
}
