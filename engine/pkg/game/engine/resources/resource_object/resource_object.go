package resource_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/pickable_object"
)

type ResourceObject struct {
	pickable_object.PickableObject
	entity.GameObject
}

func NewResourceObject(gameObj entity.IGameObject) *ResourceObject {
	resource := &ResourceObject{pickable_object.PickableObject{}, *gameObj.(*entity.GameObject)}
	resource.InitPickableObject(resource)
	return resource
}
