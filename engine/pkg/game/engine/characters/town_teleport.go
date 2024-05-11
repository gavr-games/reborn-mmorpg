package characters

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// This func is called via delayed action mechanism
// params: playerId
func TownTeleport(e entity.IEngine, params map[string]interface{}) bool {
	playerId := int(params["playerId"].(float64))
	if player, ok := e.Players().Load(playerId); ok {
		if charGameObj, charOk := e.GameObjects().Load(player.CharacterGameObjectId); charOk {
			if charGameObj.GetProperty("current_dungeon_id") != nil {
				e.SendSystemMessage("You can't teleport in dungeon.", player)
				return false
			}
			charGameObj.(entity.ICharacterObject).TownTeleport(e)

			return true
		}
	}

	return true
}