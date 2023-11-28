package melee_weapon_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

type MeleeWeaponObject struct {
	gameObj entity.IGameObject
}

func (obj *MeleeWeaponObject) InitMeleeWeaponObject(gameObj entity.IGameObject) {
	obj.gameObj = gameObj
}
