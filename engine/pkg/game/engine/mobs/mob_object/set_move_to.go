package mob_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (mob *MobObject) setMoveTo(directionChangeTime float64) {
	if targetObj, ok := mob.Engine.GameObjects()[mob.TargetObjectId]; ok {
		mob.SetMoveToCoords(&entity.MoveToCoords{
			Mode: entity.MoveToExactCoords,
			Bounds: utils.Bounds{
				X:      targetObj.X(),
				Y:      targetObj.Y(),
				Width:  0.0,
				Height: 0.0,
			},
			DirectionChangeTime: directionChangeTime,
			TimeUntilDirectionChange: 0,
		})
	}
}
