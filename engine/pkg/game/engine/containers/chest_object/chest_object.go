package bag_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/container_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/liftable_object"
)

type ChestObject struct {
	container_object.ContainerObject
	liftable_object.PickableObject
	entity.GameObject
}

func NewChestObject(gameObj entity.IGameObject) *ChestObject {
	chest := &ChestObject{container_object.ContainerObject{}, liftable_object.PickableObject{}, *gameObj.(*entity.GameObject)}
	chest.InitContainerObject(chest)
	chest.InitLiftableObject(chest)
	return chest
}
