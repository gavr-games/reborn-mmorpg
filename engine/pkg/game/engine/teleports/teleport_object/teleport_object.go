package teleport_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/building_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

type TeleportObject struct {
	building_object.BuildingObject
	entity.GameObject
}

func NewTeleportObject(gameObj entity.IGameObject) *TeleportObject {
	teleport := &TeleportObject{building_object.BuildingObject{}, *gameObj.(*entity.GameObject)}
	teleport.InitBuildingObject(teleport)
	return teleport
}
