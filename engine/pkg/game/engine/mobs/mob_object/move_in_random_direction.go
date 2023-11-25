package mob_object

import (
	"math/rand"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/constants"
)

func (mob *MobObject) moveInRandomDirection() {
	possibleDirections := constants.GetPossibleDirections()
	mobDirection := possibleDirections[rand.Intn(len(possibleDirections))]
	mob.SetXYSpeeds(mob.Engine, mobDirection)
	mob.State = MovingState
}
