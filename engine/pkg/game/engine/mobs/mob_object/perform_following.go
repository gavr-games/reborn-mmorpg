package mob_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (mob *MobObject) performFollowing(targetObj entity.IGameObject, directionChangeTime float64) {
	if mob.GetDistance(targetObj) > FollowingDistance {
		if mob.MoveToCoords() == nil {
			mob.setMoveTo(directionChangeTime)
		}
		mob.MoveToCoords().Bounds.X = targetObj.X()
		mob.MoveToCoords().Bounds.Y = targetObj.Y()
	}
	if mob.TurnToXY(targetObj.X(), targetObj.Y()) {
		mob.Engine.SendGameObjectUpdate(mob, "update_object")
	}
}
