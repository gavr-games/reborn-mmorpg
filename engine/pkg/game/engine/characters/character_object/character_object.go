package character_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/moving_object"
)

type CharacterObject struct {
	moving_object.MovingObject
	entity.GameObject
}
