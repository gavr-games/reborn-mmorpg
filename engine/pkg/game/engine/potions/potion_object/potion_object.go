package potion_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/destroyable_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/pickable_object"
)

type PotionObject struct {
	pickable_object.PickableObject
	destroyable_object.DestroyableObject
	entity.GameObject
}

func NewPotionObject(gameObj entity.IGameObject) *PotionObject {
	potion := &PotionObject{
		pickable_object.PickableObject{},
		destroyable_object.DestroyableObject{},
		*gameObj.(*entity.GameObject),
	}
	potion.InitPickableObject(potion)
	potion.InitDestroyableObject(potion)
	return potion
}
