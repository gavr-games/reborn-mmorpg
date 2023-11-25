package mob_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (mob *MobObject) performFollowing(newTickTime int64, targetObj entity.IGameObject, directionChangeTime int64) {
	if mob.GetDistance(targetObj) <= FollowingDistance {
		// Stop the mob
		if mob.Properties()["speed_x"].(float64) != 0.0 || mob.Properties()["speed_y"].(float64) != 0.0 {
			mob.Stop(mob.Engine)
		}
		mob.turnToTarget(targetObj) // TODO: send only on change
	} else {
		if (newTickTime - mob.directionTickTime >= directionChangeTime) {
			mob.directionTickTime = newTickTime
			mob.SetXYSpeeds(mob.Engine, mob.getDirectionToTarget(targetObj))
		}
	}
}
