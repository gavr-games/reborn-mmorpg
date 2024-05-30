package evolvable_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

type EvolvableObject struct {
	gameObj entity.IGameObject
}

func (obj *EvolvableObject) InitEvolvableObject(gameObj entity.IGameObject) {
	obj.gameObj = gameObj
}
