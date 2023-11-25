package mob_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (mob *MobObject) turnToTarget(targetObj entity.IGameObject) {
	direction := mob.getDirectionToTarget(targetObj)
	if mob.Rotation() != mob.GetRotationByDirection(direction) {
		mob.SetRotationByDirection(direction)
		mob.Engine.SendGameObjectUpdate(mob, "update_object")
	}
}
