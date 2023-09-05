package engine

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

// send new state of the game object to all players who can see it
func SendGameObjectUpdate(e IEngine, gameObj *entity.GameObject, updateType string) {
	SendResponseToVisionAreas(e, gameObj, updateType, map[string]interface{}{
		"object": gameObj,
	})
	storage.GetClient().Updates <- gameObj
}
