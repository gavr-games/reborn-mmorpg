package delayed_actions

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects"
)

func Cancel(e entity.IEngine, gameObj *entity.GameObject) bool {
	if gameObj.CurrentAction == nil {
		return true
	}

	delayedActionFuncName := gameObj.CurrentAction.FuncName

	gameObj.CurrentAction = nil

	storage.GetClient().Updates <- gameObj

	e.SendResponseToVisionAreas(gameObj, "cancel_delayed_action", map[string]interface{}{
		"object": game_objects.GetInfo(e.GameObjects(), gameObj),
		"action": delayedActionFuncName,
	})

	return true
}