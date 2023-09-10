package game_objects

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// Intersect - Checks if a game object intersects with another game object
func Intersect(a *entity.GameObject, b *entity.GameObject) bool {

	aMaxX := a.X + a.Width
	aMaxY := a.Y + a.Height
	bMaxX := b.X + b.Width
	bMaxY := b.Y + b.Height

	// a is left of b
	if aMaxX < b.X {
		return false
	}

	// a is right of b
	if a.X > bMaxX {
		return false
	}

	// a is above b
	if aMaxY < b.Y {
		return false
	}

	// a is below b
	if a.Y > bMaxY {
		return false
	}

	// The two overlap
	return true
}