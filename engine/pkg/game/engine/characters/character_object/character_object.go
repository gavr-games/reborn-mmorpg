package character_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/moving_object"
)

type CharacterObject struct {
	moving_object.MovingObject
	entity.GameObject
}

func NewCharacterObject(e entity.IEngine, gameObj entity.IGameObject) *CharacterObject {
	character := &CharacterObject{moving_object.MovingObject{}, *gameObj.(*entity.GameObject)}
	character.InitMovingObject(character)
	return character
}
