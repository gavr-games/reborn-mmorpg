package weapon_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/destroyable_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/pickable_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/equipable_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/melee_weapon_object"
)

type WeaponObject struct {
	pickable_object.PickableObject
	equipable_object.EquipableObject
	melee_weapon_object.MeleeWeaponObject
	destroyable_object.DestroyableObject
	entity.GameObject
}

func NewWeaponObject(gameObj entity.IGameObject) *WeaponObject {
	weapon := &WeaponObject{
		pickable_object.PickableObject{},
		equipable_object.EquipableObject{},
		melee_weapon_object.MeleeWeaponObject{},
		destroyable_object.DestroyableObject{},
		*gameObj.(*entity.GameObject),
	}
	weapon.InitPickableObject(weapon)
	weapon.InitEquipableObject(weapon)
	weapon.InitMeleeWeaponObject(weapon)
	weapon.InitDestroyableObject(weapon)
	return weapon
}
