package characters

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/constants"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/serializers"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/targets"
)

func Reborn(e entity.IEngine, charGameObj *entity.GameObject) {
	charGameObj.Properties["health"] = charGameObj.Properties["max_health"]
	targets.Deselect(e, charGameObj)
	
	// Cancel delayed action
	// TODO: refactor code so we can reuse delayed_actions.Cancel
	if charGameObj.CurrentAction != nil {
		delayedActionFuncName := charGameObj.CurrentAction.FuncName

		charGameObj.CurrentAction = nil

		storage.GetClient().Updates <- game_objects.Clone(charGameObj)

		e.SendResponseToVisionAreas(charGameObj, "cancel_delayed_action", map[string]interface{}{
			"object": serializers.GetInfo(e.GameObjects(), charGameObj),
			"action": delayedActionFuncName,
		})
	}

	Move(e, charGameObj, constants.InitialPlayerX, constants.InitialPlayerY)
	e.SendGameObjectUpdate(charGameObj, "update_object")
}