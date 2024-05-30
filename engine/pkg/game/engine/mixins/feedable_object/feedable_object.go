package feedable_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

type FeedableObject struct {
	gameObj entity.IGameObject
}

func (obj *FeedableObject) InitFeedableObject(gameObj entity.IGameObject) {
	obj.gameObj = gameObj
}
