package hatchery_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/building_object"
)

type HatcheryObject struct {
	building_object.BuildingObject
	entity.GameObject
}

func NewHatcheryObject(gameObj entity.IGameObject) *HatcheryObject {
	hatchery := &HatcheryObject{building_object.BuildingObject{}, *gameObj.(*entity.GameObject)}
	hatchery.InitBuildingObject(hatchery)
	return hatchery
}
