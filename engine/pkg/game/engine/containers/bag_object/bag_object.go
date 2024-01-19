package bag_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/container_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/pickable_object"
)

type BagObject struct {
	container_object.ContainerObject
	pickable_object.PickableObject
	entity.GameObject
}

func NewBagObject(gameObj entity.IGameObject) *BagObject {
	bag := &BagObject{container_object.ContainerObject{}, pickable_object.PickableObject{}, *gameObj.(*entity.GameObject)}
	bag.InitContainerObject(bag)
	bag.InitPickableObject(bag)
	return bag
}
