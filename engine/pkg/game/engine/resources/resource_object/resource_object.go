package resource_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/destroyable_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/pickable_object"
)

type ResourceObject struct {
	pickable_object.PickableObject
	destroyable_object.DestroyableObject
	entity.GameObject
}

func NewResourceObject(gameObj entity.IGameObject) *ResourceObject {
	resource := &ResourceObject{
		pickable_object.PickableObject{},
		destroyable_object.DestroyableObject{},
		*gameObj.(*entity.GameObject),
	}
	resource.InitPickableObject(resource)
	resource.InitDestroyableObject(resource)
	return resource
}
