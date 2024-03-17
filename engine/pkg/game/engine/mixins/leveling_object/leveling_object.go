package leveling_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

type LevelingObject struct {
	gameObj entity.IGameObject
}

func (obj *LevelingObject) InitLevelingObject(gameObj entity.IGameObject) {
	obj.gameObj = gameObj
}
