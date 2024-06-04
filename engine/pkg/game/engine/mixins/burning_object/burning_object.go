package burning_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

type BurningObject struct {
	gameObj entity.IGameObject
}

func (obj *BurningObject) InitBurningObject(gameObj entity.IGameObject) {
	obj.gameObj = gameObj
}
