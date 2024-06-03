package door_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/building_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

type DoorObject struct {
	building_object.BuildingObject
	entity.GameObject
}

func NewDoorObject(gameObj entity.IGameObject) *DoorObject {
	door := &DoorObject{building_object.BuildingObject{}, *gameObj.(*entity.GameObject)}
	door.InitBuildingObject(door)
	return door
}
