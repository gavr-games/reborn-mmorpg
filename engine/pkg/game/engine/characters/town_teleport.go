package characters

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/constants"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// This func is called via delayed action mechanism
// params: playerId
func TownTeleport(e entity.IEngine, params map[string]interface{}) bool {
	playerId := int(params["playerId"].(float64))
	if player, ok := e.Players().Load(playerId); ok {
		if charGameObj, charOk := e.GameObjects().Load(player.CharacterGameObjectId); charOk {
			charGameObj.(entity.ICharacterObject).DeselectTarget(e)
			charGameObjClone := charGameObj.Clone()
			e.SendResponseToVisionAreas(charGameObjClone, "remove_object", map[string]interface{}{
				"object": charGameObjClone,
			})
			charGameObj.(entity.ICharacterObject).Move(e, constants.InitialPlayerX, constants.InitialPlayerY, 0)
			e.SendGameObjectUpdate(charGameObj, "update_object")
		}
	}

	return true
}