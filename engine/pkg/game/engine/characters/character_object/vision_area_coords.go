package character_object

import (
	"math"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/constants"
)

func (charGameObj *CharacterObject) GetVisionAreaX() float64 {
	return math.Ceil(charGameObj.X() - constants.PlayerVisionArea / 2 - 5.0)
}

func (charGameObj *CharacterObject) GetVisionAreaY() float64 {
	return math.Ceil(charGameObj.Y() - constants.PlayerVisionArea / 2 + 5.0)
}
