package delayed_actions

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/serializers"
)

func Start(e entity.IEngine, gameObj *entity.GameObject, funcName string, params map[string]interface{}) bool {
	delayedAction := &entity.DelayedAction{
		FuncName: funcName,
		Params: params,
		TimeLeft: GetDelayedActionsAtlas()[funcName]["duration"].(float64),
	}

	gameObj.CurrentAction = delayedAction

	storage.GetClient().Updates <- game_objects.Clone(gameObj)

	e.SendResponseToVisionAreas(gameObj, "start_delayed_action", map[string]interface{}{
		"object": serializers.GetInfo(e.GameObjects(), gameObj),
		"duration": delayedAction.TimeLeft,
		"action": funcName,
	})

	return true
}
