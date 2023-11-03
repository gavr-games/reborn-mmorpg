package engine

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/characters"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects"
)

// move players
func MovePlayers(e entity.IEngine, tickDelta int64) {
	for _, player := range e.Players() {
    if player.Client != nil && player.CharacterGameObjectId != "" && player.VisionAreaGameObjectId != "" {
			charGameObj := e.GameObjects()[player.CharacterGameObjectId]
			speedX := charGameObj.Properties["speed_x"].(float64)
			speedY := charGameObj.Properties["speed_y"].(float64)
			if speedX != 0 || speedY != 0 {
				dx := speedX / 1000.0 * float64(tickDelta)
				dy := speedY / 1000.0 * float64(tickDelta)

				dx, dy = game_objects.CanMove(e.Floors()[charGameObj.Floor], charGameObj, dx, dy)

				// Stop the object
				if dx == 0.0 && dy == 0.0 {
					charGameObj.Properties["speed_x"] = 0.0
					charGameObj.Properties["speed_y"] = 0.0
					e.SendGameObjectUpdate(charGameObj, "update_object")
					continue
				}

				// Update player character game object
				characters.Move(e, charGameObj, charGameObj.X + dx, charGameObj.Y + dy)
			}
		}
	}
}
