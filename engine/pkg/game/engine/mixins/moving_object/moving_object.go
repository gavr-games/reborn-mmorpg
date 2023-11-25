package moving_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

type MovingObject struct {
	gameObj entity.IGameObject
}

func (mObj *MovingObject) InitMovingObject(gameObj entity.IGameObject) {
	mObj.gameObj = gameObj
}
