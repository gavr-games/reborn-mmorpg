package melee_weapon_object

import (
	"math"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

type WorldPoint struct {
	X, Y float64
}

// Checks 21 points inside and on edges of weapon hit sector
// (you can see this sector on frontend highlighted with red during melee hit)
// If they are inside the target obj (rectangle or circle) - then CanHit -> true
func (obj *MeleeWeaponObject) CanHit(attacker entity.IGameObject, target entity.IGameObject) bool {
	weapon := obj.gameObj
	points := make([]WorldPoint, 21)

	// Add center point
	points[0] = WorldPoint{attacker.X(), attacker.Y()}

	// Add points for 5 vectors of the sector
	hitAngle := weapon.Properties()["hit_angle"].(float64) * (math.Pi / 180.0)
	startingAngle := attacker.Rotation() - hitAngle / 2.0
	stepAngle := hitAngle / 4.0
	stepRadius := weapon.Properties()["hit_radius"].(float64) / 4

	pointIndex := 1

	for i := 0; i < 5; i++ { // 5 vectors from attacker point to sector arc edge
		angle := startingAngle + stepAngle * float64(i)
		for k := 1; k < 5; k++ { // take 4 points on each vector, except center (added earlier)
			radius := stepRadius * float64(k)
			points[pointIndex] = WorldPoint{attacker.X() + radius * math.Cos(angle), attacker.Y() + radius * math.Sin(angle)}
			pointIndex++
		}
	}
	// check points are inside the target shape
	// if at least one point is inside then the weapon can hit the target
	for p := 0; p < len(points); p++ {
		point := points[p]
		if target.Properties()["shape"].(string) == "rectangle" {
			if point.X >= target.X() && point.X <= target.X() + target.Width() &&
					point.Y >= target.Y() && point.Y <= target.Y() + target.Height() {
						return true
					}
		} else
		if target.Properties()["shape"].(string) == "circle" {
			// distance between points <= target radius
			if math.Pow(point.X - target.X(), 2.0) + math.Pow(point.Y - target.Y(), 2.0) <= math.Pow(target.Width() / 2.0, 2.0) {
				return true
			}
		}
	}

	return false
}
