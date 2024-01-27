package chest_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/destroyable_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/container_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/liftable_object"
)

type ChestObject struct {
	container_object.ContainerObject
	liftable_object.LiftableObject
	destroyable_object.DestroyableObject
	entity.GameObject
}

func NewChestObject(gameObj entity.IGameObject) *ChestObject {
	chest := &ChestObject{
		container_object.ContainerObject{},
		liftable_object.LiftableObject{},
		destroyable_object.DestroyableObject{},
		*gameObj.(*entity.GameObject),
	}
	chest.InitContainerObject(chest)
	chest.InitLiftableObject(chest)
	chest.InitDestroyableObject(chest)
	return chest
}
