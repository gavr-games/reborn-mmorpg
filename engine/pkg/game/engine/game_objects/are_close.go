package game_objects

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

const (
	MaxDistance = 0.1
)

// Determines if 2 objects are close enough to each other
func AreClose(a *entity.GameObject, b *entity.GameObject) bool {
	if (a.Floor != b.Floor) {
		return false
	}
	return GetDistance(a, b) < MaxDistance
}
