package pickable_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

type PickableObject struct {
	gameObj entity.IGameObject
}

func (obj *PickableObject) InitPickableObject(gameObj entity.IGameObject) {
	obj.gameObj = gameObj
}
