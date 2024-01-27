package bag_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/destroyable_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/container_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/pickable_object"
)

type BagObject struct {
	container_object.ContainerObject
	pickable_object.PickableObject
	destroyable_object.DestroyableObject
	entity.GameObject
}

func NewBagObject(gameObj entity.IGameObject) *BagObject {
	bag := &BagObject{
		container_object.ContainerObject{},
		pickable_object.PickableObject{},
		destroyable_object.DestroyableObject{},
		*gameObj.(*entity.GameObject),
	}
	bag.InitContainerObject(bag)
	bag.InitDestroyableObject(bag)
	bag.InitPickableObject(bag)
	return bag
}
