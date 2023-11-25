package container_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

type ContainerObject struct {
	gameObj entity.IGameObject
}

func (cObj *ContainerObject) InitContainerObject(gameObj entity.IGameObject) {
	cObj.gameObj = gameObj
}
