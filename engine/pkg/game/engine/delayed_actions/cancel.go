package delayed_actions

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/serializers"
)

func Cancel(e entity.IEngine, gameObj entity.IGameObject) bool {
	if gameObj.CurrentAction() == nil {
		return true
	}

	delayedActionFuncName := gameObj.CurrentAction().FuncName

	gameObj.SetCurrentAction(nil)

	storage.GetClient().Updates <- gameObj.Clone()

	e.SendResponseToVisionAreas(gameObj, "cancel_delayed_action", map[string]interface{}{
		"object": serializers.GetInfo(e, gameObj),
		"action": delayedActionFuncName,
	})

	return true
}
