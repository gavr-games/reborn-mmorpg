package mobs

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// move mobs and trigger mob logic
func Update(e entity.IEngine, tickDelta int64, newTickTime int64) {
	for _, mob := range e.Mobs() {
		mobObj := mob.(entity.IGameObject)
    if mobObj != nil {
			// Trigger Mob logic
			mob.Run(newTickTime)

			// Trigger Move to Coords logic
			mobObj.(entity.IMovingObject).PerformMoveTo(e, tickDelta)

			// Move mobs
			speedX := mobObj.Properties()["speed_x"].(float64)
			speedY := mobObj.Properties()["speed_y"].(float64)
			if speedX != 0 || speedY != 0 {
				dx := speedX / 1000.0 * float64(tickDelta)
				dy := speedY / 1000.0 * float64(tickDelta)

				dx, dy = mobObj.(entity.IMovingObject).CanMove(e, dx, dy, false)

				// Stop the object
				if dx == 0.0 && dy == 0.0 {
					mobObj.(entity.IMovingObject).Stop(e)
					continue
				}

				// Update mob game object
				e.Floors()[mobObj.Floor()].FilteredRemove(mobObj, func(b utils.IBounds) bool {
					return mobObj.Id() == b.(entity.IGameObject).Id()
				})
				mobObj.SetX(mobObj.X() + dx)
				mobObj.SetY(mobObj.Y() + dy)
				e.Floors()[mobObj.Floor()].Insert(mobObj)
			}
		}
	}
}
