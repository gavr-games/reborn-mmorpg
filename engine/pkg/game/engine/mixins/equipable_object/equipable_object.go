package equipable_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

type EquipableObject struct {
	gameObj entity.IGameObject
}

func (obj *EquipableObject) InitEquipableObject(gameObj entity.IGameObject) {
	obj.gameObj = gameObj
}
