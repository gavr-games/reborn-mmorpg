package delayed_actions

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/serializers"
)

func Start(e entity.IEngine, gameObj entity.IGameObject, funcName string, params map[string]interface{}, timeLeft float64) bool {
	if timeLeft == -1.0 {
		timeLeft = GetDelayedActionsAtlas()[funcName]["duration"].(float64)
	}
	delayedAction := &entity.DelayedAction{
		FuncName: funcName,
		Params: params,
		TimeLeft: timeLeft,
	}

	gameObj.SetCurrentAction(delayedAction)

	storage.GetClient().Updates <- gameObj.Clone()

	e.SendResponseToVisionAreas(gameObj, "start_delayed_action", map[string]interface{}{
		"object": serializers.GetInfo(e.GameObjects(), gameObj),
		"duration": delayedAction.TimeLeft,
		"action": funcName,
	})

	return true
}
