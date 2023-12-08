package delayed_actions

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

func Start(e entity.IEngine, gameObj entity.IGameObject, funcName string, params map[string]interface{}, timeLeft float64) bool {
	if timeLeft == -1.0 {
		timeLeft = GetDelayedActionsAtlas()[funcName]["duration"].(float64)
	}
	delayedAction := &entity.DelayedAction{
		FuncName: funcName,
		Params: params,
		TimeLeft: timeLeft,
		Status: entity.DelayedActionReady,
	}

	gameObj.SetCurrentAction(delayedAction)

	storage.GetClient().Updates <- gameObj.Clone()

	return true
}
