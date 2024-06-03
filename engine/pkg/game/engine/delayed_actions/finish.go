package delayed_actions

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/serializers"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

func Finish(e entity.IEngine, gameObj entity.IGameObject) bool {
	currentAction := gameObj.CurrentAction()
	if currentAction == nil {
		return true
	}

	delayedActionFuncName := currentAction.FuncName()

	// Call delayed function
	// all delayed fucntions must be func(entity.IEngine, map[string]interface{}) bool
	delayedFunc := GetDelayedActionsAtlas()[delayedActionFuncName]["func"].(func(entity.IEngine, map[string]interface{}) bool)

	gameObj.SetCurrentAction(nil)
	e.DelayedActions().Delete(gameObj.Id())

	storage.GetClient().Updates <- gameObj.Clone()

	delayedFunc(e, currentAction.Params())

	e.SendResponseToVisionAreas(gameObj, "finish_delayed_action", map[string]interface{}{
		"object": serializers.GetInfo(e, gameObj),
		"action": delayedActionFuncName,
	})

	return true
}
