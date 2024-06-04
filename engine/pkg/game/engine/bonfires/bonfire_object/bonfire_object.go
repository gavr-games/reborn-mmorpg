package bonfire_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/building_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/burning_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

type BonfireObject struct {
	burning_object.BurningObject
	building_object.BuildingObject
	entity.GameObject
}

func NewBonfireObject(gameObj entity.IGameObject) *BonfireObject {
	bonfire := &BonfireObject{burning_object.BurningObject{}, building_object.BuildingObject{}, *gameObj.(*entity.GameObject)}
	bonfire.InitBurningObject(bonfire)
	bonfire.InitBuildingObject(bonfire)
	return bonfire
}
