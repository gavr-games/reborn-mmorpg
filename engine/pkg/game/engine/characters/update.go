package characters

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// Move characters
func Update(e entity.IEngine, tickDelta int64) {
	for _, player := range e.Players() {
    if player.Client != nil && player.CharacterGameObjectId != "" && player.VisionAreaGameObjectId != "" {
			charGameObj := e.GameObjects()[player.CharacterGameObjectId]
			speedX := charGameObj.Properties()["speed_x"].(float64)
			speedY := charGameObj.Properties()["speed_y"].(float64)
			if speedX != 0 || speedY != 0 {
				dx := speedX / 1000.0 * float64(tickDelta)
				dy := speedY / 1000.0 * float64(tickDelta)

				dx, dy = charGameObj.(entity.IMovingObject).CanMove(e, dx, dy)

				// Stop the object
				if dx == 0.0 && dy == 0.0 {
					charGameObj.(entity.IMovingObject).Stop(e)
					continue
				}

				// Update player character game object
				charGameObj.(entity.ICharacterObject).Move(e, charGameObj.X() + dx, charGameObj.Y() + dy)
			}
		}
	}
}
