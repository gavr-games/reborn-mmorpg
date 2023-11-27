package wall_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/building_object"
)

type WallObject struct {
	building_object.BuildingObject
	entity.GameObject
}

func NewWallObject(gameObj entity.IGameObject) *WallObject {
	wall := &WallObject{building_object.BuildingObject{}, *gameObj.(*entity.GameObject)}
	wall.InitBuildingObject(wall)
	return wall
}
