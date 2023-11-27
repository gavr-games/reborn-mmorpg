package building_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

type BuildingObject struct {
	gameObj entity.IGameObject
}

func (bObj *BuildingObject) InitBuildingObject(gameObj entity.IGameObject) {
	bObj.gameObj = gameObj
}
