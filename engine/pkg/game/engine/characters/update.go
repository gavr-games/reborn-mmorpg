package characters

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// Move characters
func Update(e entity.IEngine, tickDelta int64) {
	e.Players().Range(func(playerId int, player *entity.Player) bool {
    if player.Client != nil && player.CharacterGameObjectId != "" && player.VisionAreaGameObjectId != "" {
			charGameObj := e.GameObjects()[player.CharacterGameObjectId]

			// Trigger Move to Coords logic
			charGameObj.(entity.IMovingObject).PerformMoveTo(e, tickDelta)

			speedX := charGameObj.Properties()["speed_x"].(float64)
			speedY := charGameObj.Properties()["speed_y"].(float64)
			if speedX != 0 || speedY != 0 {
				dx := speedX / 1000.0 * float64(tickDelta)
				dy := speedY / 1000.0 * float64(tickDelta)

				dx, dy = charGameObj.(entity.IMovingObject).CanMove(e, dx, dy, false)

				// Stop the object
				if dx == 0.0 && dy == 0.0 {
					charGameObj.(entity.IMovingObject).Stop(e)
					return true
				}

				// Update player character game object
				charGameObj.(entity.ICharacterObject).Move(e, charGameObj.X() + dx, charGameObj.Y() + dy)
			}
		}
		return true
	})
}
