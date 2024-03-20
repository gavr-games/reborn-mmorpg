package character_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/constants"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/serializers"
)

func (charGameObj *CharacterObject) Reborn(e entity.IEngine) {
	charGameObj.SetProperty("health", charGameObj.GetProperty("max_health"))
	charGameObj.DeselectTarget(e)
	
	// Cancel delayed action
	// TODO: refactor code so we can reuse delayed_actions.Cancel
	if charGameObj.CurrentAction() != nil {
		delayedActionFuncName := charGameObj.CurrentAction().FuncName()

		charGameObj.SetCurrentAction(nil)

		storage.GetClient().Updates <- charGameObj.Clone()

		e.SendResponseToVisionAreas(charGameObj, "cancel_delayed_action", map[string]interface{}{
			"object": serializers.GetInfo(e, charGameObj),
			"action": delayedActionFuncName,
		})
	}
	charGameObjClone := charGameObj.Clone()
	e.SendResponseToVisionAreas(charGameObjClone, "remove_object", map[string]interface{}{
		"object": charGameObjClone,
	})
	charGameObj.Move(e, constants.InitialPlayerX, constants.InitialPlayerY, 0)
	e.SendGameObjectUpdate(charGameObj, "update_object")
}
