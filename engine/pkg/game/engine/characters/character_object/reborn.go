package character_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/constants"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/dungeons"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/serializers"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

func (charGameObj *CharacterObject) Reborn(e entity.IEngine) {
	charGameObj.SetProperty("health", charGameObj.GetProperty("max_health"))
	charGameObj.SetProperty("last_death", float64(e.CurrentTickTime()))
	charGameObj.DeselectTarget(e)
	
	// Cancel delayed action
	// TODO: refactor code so we can reuse delayed_actions.Cancel
	if charGameObj.CurrentAction() != nil {
		delayedActionFuncName := charGameObj.CurrentAction().FuncName()

		charGameObj.SetCurrentAction(nil)
		e.DelayedActions().Delete(charGameObj.Id())

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
	charGameObj.Move(e, constants.InitialPlayerX, constants.InitialPlayerY, e.GetGameAreaByKey(constants.InitialPlayerArea).Id())
	e.SendGameObjectUpdate(charGameObj, "update_object")
	
	// Destroy the dungeon if char died in dungeon
	go dungeons.Destroy(e, charGameObj)
}
