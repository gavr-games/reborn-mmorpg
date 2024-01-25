package liftable_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

type LiftableObject struct {
	gameObj entity.IGameObject
}

func (obj *LiftableObject) InitLiftableObject(gameObj entity.IGameObject) {
	obj.gameObj = gameObj
}
