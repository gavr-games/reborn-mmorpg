package backpack_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/container_object"
)

type BackpackObject struct {
	container_object.ContainerObject
	entity.GameObject
}

func NewBackpackObject(gameObj entity.IGameObject) *BackpackObject {
	backpack := &BackpackObject{container_object.ContainerObject{}, *gameObj.(*entity.GameObject)}
	backpack.InitContainerObject(backpack)
	return backpack
}
