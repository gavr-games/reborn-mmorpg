package delayed_actions

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/craft"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/serializers"
)

// Difference from usual Start is that we take TimeLeft from craft atlas
func StartCraft(e entity.IEngine, gameObj *entity.GameObject, funcName string, params map[string]interface{}) bool {
	craftItem := params["item_name"].(string)
	delayedAction := &entity.DelayedAction{
		FuncName: funcName,
		Params: params,
		TimeLeft: craft.GetAtlas()[craftItem].(map[string]interface{})["duration"].(float64),
	}

	gameObj.CurrentAction = delayedAction

	storage.GetClient().Updates <- gameObj

	e.SendResponseToVisionAreas(gameObj, "start_delayed_action", map[string]interface{}{
		"object": serializers.GetInfo(e.GameObjects(), gameObj),
		"duration": delayedAction.TimeLeft,
		"action": funcName,
	})

	return true
}
