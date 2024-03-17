package character_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/leveling_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/moving_object"
)

type CharacterObject struct {
	moving_object.MovingObject
	leveling_object.LevelingObject
	entity.GameObject
}

func NewCharacterObject(gameObj entity.IGameObject) *CharacterObject {
	character := &CharacterObject{
		moving_object.MovingObject{},
		leveling_object.LevelingObject{},
		*gameObj.(*entity.GameObject),
	}
	character.InitMovingObject(character)
	character.InitLevelingObject(character)
	return character
}
