package game_objects

import (
	"math"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// Get approximate distance between objects. Assuming all of them are rectangles
func GetDistance(a *entity.GameObject, b *entity.GameObject) float64 {
	aXCenter := a.X + a.Width / 2
	aYCenter := a.Y + a.Height / 2
	
	bXCenter := b.X + b.Width / 2
	bYCenter := b.Y + b.Height / 2

	xDistance := math.Abs(aXCenter - bXCenter) - (a.Width / 2 + b.Width / 2)
	if xDistance < 0 {
		xDistance = 0.0
	}

	yDistance := math.Abs(aYCenter - bYCenter) - (a.Height / 2 + b.Height / 2)
	if yDistance < 0 {
		yDistance = 0.0
	}

	return math.Sqrt(math.Pow(xDistance, 2.0) + math.Pow(yDistance, 2.0))
}