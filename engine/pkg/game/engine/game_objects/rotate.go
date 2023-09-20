package game_objects

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// Rotates Game object. Possible rotations 0 and 1 (0 and 90 dergrees)
func Rotate(gameObj *entity.GameObject, rotation float64) {
	if gameObj.Rotation != rotation {
		gameObj.Rotation = rotation
		width := gameObj.Width
		gameObj.Width = gameObj.Height
		gameObj.Height = width
	}
}