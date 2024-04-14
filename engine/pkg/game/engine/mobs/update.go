package mobs

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// move mobs and trigger mob logic
func Update(e entity.IEngine, tickDelta int64, newTickTime int64) {
	e.Mobs().Range(func(id string, mob entity.IMobObject) bool {
		mobObj := mob.(entity.IGameObject)
    if mobObj != nil {
			// Trigger Mob logic
			mob.Run(newTickTime)

			// Trigger Move to Coords logic
			mobObj.(entity.IMovingObject).PerformMoveTo(e, tickDelta)

			// Move mobs
			speedX := mobObj.GetProperty("speed_x").(float64)
			speedY := mobObj.GetProperty("speed_y").(float64)
			if speedX != 0 || speedY != 0 {
				dx := speedX / 1000.0 * float64(tickDelta)
				dy := speedY / 1000.0 * float64(tickDelta)

				dx, dy = mobObj.(entity.IMovingObject).CanMove(e, dx, dy, false)

				// Stop the object
				if dx == 0.0 && dy == 0.0 {
					mobObj.(entity.IMovingObject).Stop(e)
					return true
				}

				// Update mob game object
				gameArea, gaOk := e.GameAreas().Load(mobObj.GameAreaId())
				if gaOk {
					gameArea.FilteredRemove(mobObj, func(b utils.IBounds) bool {
						return mobObj.Id() == b.(entity.IGameObject).Id()
					})
				}
				mobObj.SetX(mobObj.X() + dx)
				mobObj.SetY(mobObj.Y() + dy)
				gameArea.Insert(mobObj)
			}
		}
		return true
	})
}
