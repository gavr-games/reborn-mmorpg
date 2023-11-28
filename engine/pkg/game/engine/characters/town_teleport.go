package characters

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/constants"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// This func is called via delayed action mechanism
// params: playerId
func TownTeleport(e entity.IEngine, params map[string]interface{}) bool {
	playerId := int(params["playerId"].(float64))
	if player, ok := e.Players()[playerId]; ok {
		charGameObj := e.GameObjects()[player.CharacterGameObjectId]
		charGameObj.(entity.ICharacterObject).DeselectTarget(e)
		charGameObj.(entity.ICharacterObject).Move(e, constants.InitialPlayerX, constants.InitialPlayerY)
		e.SendGameObjectUpdate(charGameObj, "update_object")
	}

	return true
}