package delayed_actions

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects"
)

func Finish(e entity.IEngine, gameObj *entity.GameObject) bool {
	if gameObj.CurrentAction == nil {
		return true
	}

	delayedActionFuncName := gameObj.CurrentAction.FuncName

	// Call delayed function
	// all delayed fucntions must be func(entity.IEngine, map[string]interface{}) bool
	delayedFunc := GetDelayedActionsAtlas()[delayedActionFuncName]["func"].(func(entity.IEngine, map[string]interface{}) bool)
	delayedFunc(e, gameObj.CurrentAction.Params)

	gameObj.CurrentAction = nil

	storage.GetClient().Updates <- gameObj

	e.SendResponseToVisionAreas(gameObj, "finish_delayed_action", map[string]interface{}{
		"object": game_objects.GetInfo(e.GameObjects(), gameObj),
		"action": delayedActionFuncName,
	})

	return true
}
