package mob_object

import (
	"math"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/constants"
)

func (mob *MobObject) getDirectionToTarget(targetObj entity.IGameObject) string {
	possibleDirections := constants.GetPossibleDirections()
	// Calclate angle between mob and target
	// Choose the closest direction by angle by calculatin index in PossibleDirections slice
	dx := targetObj.X() - mob.X()
	dy := targetObj.Y() - mob.Y()
	angle := math.Atan2(dy, dx) // range (-PI, PI)
	if angle < 0.0 {
		angle = angle + math.Pi * 2
	}
	quotient := math.Floor(angle / (math.Pi / 4)) // math.Pi / 4 - is the angle between movement directions
	remainder := angle - (math.Pi / 4) * quotient
	if (remainder > math.Pi / 8) {
		quotient = quotient + 1.0
	}
	directionIndex := int(quotient)
	if (directionIndex == len(possibleDirections)) {
		directionIndex = 0
	}
	return possibleDirections[directionIndex]
}
