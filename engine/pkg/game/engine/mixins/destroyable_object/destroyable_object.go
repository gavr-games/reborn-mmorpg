package destroyable_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

type DestroyableObject struct {
	gameObj entity.IGameObject
}

func (obj *DestroyableObject) InitDestroyableObject(gameObj entity.IGameObject) {
	obj.gameObj = gameObj
}
