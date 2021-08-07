package engine

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// move players
func MovePlayers(e IEngine, tickDelta int64) {
	for _, player := range e.Players() {
    if player.Client != nil && player.CharacterGameObjectId != "" && player.VisionAreaGameObjectId != "" {
			charGameObj := e.GameObjects()[player.CharacterGameObjectId]
			visionAreaGameObj := e.GameObjects()[player.VisionAreaGameObjectId]
			speedX := charGameObj.Properties["speed_x"].(float64)
			speedY := charGameObj.Properties["speed_y"].(float64)
			if speedX != 0 || speedY != 0 {
				// Update player character game object
				e.Floors()[charGameObj.Floor].FilteredRemove(charGameObj, func(b utils.IBounds) bool {
					return charGameObj.Id == b.(*entity.GameObject).Id
				})
				charGameObj.X = charGameObj.X + speedX / 1000.0 * float64(tickDelta)
				charGameObj.Y = charGameObj.Y + speedY / 1000.0 * float64(tickDelta)
				charGameObj.Properties["x"] = charGameObj.X
				charGameObj.Properties["y"] = charGameObj.Y
				e.Floors()[charGameObj.Floor].Insert(charGameObj)
				// Update vision area game object
				e.Floors()[visionAreaGameObj.Floor].FilteredRemove(visionAreaGameObj, func(b utils.IBounds) bool {
					return visionAreaGameObj.Id == b.(*entity.GameObject).Id
				})
				visionAreaGameObj.X = visionAreaGameObj.X + speedX / 1000.0 * float64(tickDelta)
				visionAreaGameObj.Y = visionAreaGameObj.Y + speedY / 1000.0 * float64(tickDelta)
				visionAreaGameObj.Properties["x"] = visionAreaGameObj.X
				visionAreaGameObj.Properties["y"] = visionAreaGameObj.Y
				e.Floors()[visionAreaGameObj.Floor].Insert(visionAreaGameObj)

				// determine new and old visible objects
				// send remove old visible objects
				// send add new visible objects
			}
		}
	}
}
