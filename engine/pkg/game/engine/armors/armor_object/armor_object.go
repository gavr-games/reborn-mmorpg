package armor_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/destroyable_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/pickable_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/equipable_object"
)

type ArmorObject struct {
	pickable_object.PickableObject
	equipable_object.EquipableObject
	destroyable_object.DestroyableObject
	entity.GameObject
}

func NewArmorObject(gameObj entity.IGameObject) *ArmorObject {
	armor := &ArmorObject{
		pickable_object.PickableObject{},
		equipable_object.EquipableObject{},
		destroyable_object.DestroyableObject{},
		*gameObj.(*entity.GameObject),
	}
	armor.InitPickableObject(armor)
	armor.InitEquipableObject(armor)
	armor.InitDestroyableObject(armor)
	return armor
}
