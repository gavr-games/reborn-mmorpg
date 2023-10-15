package mobs

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects"
)

// move mobs and trigger mob logic
func Update(e entity.IEngine, tickDelta int64, newTickTime int64) {
	for _, mob := range e.Mobs() {
		mobObj := e.GameObjects()[mob.GetId()]

    if mobObj != nil {
			// Trigger Mob logic
			mob.Run(newTickTime)

			// Move mobs
			speedX := mobObj.Properties["speed_x"].(float64)
			speedY := mobObj.Properties["speed_y"].(float64)
			if speedX != 0 || speedY != 0 {
				dx := speedX / 1000.0 * float64(tickDelta)
				dy := speedY / 1000.0 * float64(tickDelta)

				dx, dy = game_objects.CanMove(e.Floors()[mobObj.Floor], mobObj, dx, dy)

				// Stop the object
				if dx == 0.0 && dy == 0.0 {
					mobObj.Properties["speed_x"] = 0.0
					mobObj.Properties["speed_y"] = 0.0
					e.SendGameObjectUpdate(mobObj, "update_object")
					continue
				}

				// Update mob game object
				e.Floors()[mobObj.Floor].FilteredRemove(mobObj, func(b utils.IBounds) bool {
					return mobObj.Id == b.(*entity.GameObject).Id
				})
				mobObj.X += dx
				mobObj.Y += dy
				mobObj.Properties["x"] = mobObj.Properties["x"].(float64) + dx
				mobObj.Properties["y"] = mobObj.Properties["y"].(float64) + dy
				e.Floors()[mobObj.Floor].Insert(mobObj)
			}
		}
	}
}
