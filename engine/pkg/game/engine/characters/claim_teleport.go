package characters

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/targets"
)

// This func is called via delayed action mechanism
// params: playerId
func ClaimTeleport(e entity.IEngine, params map[string]interface{}) bool {
	playerId := int(params["playerId"].(float64))
	if player, ok := e.Players()[playerId]; ok {
		charGameObj := e.GameObjects()[player.CharacterGameObjectId]
		targets.Deselect(e, charGameObj)
		obelisk := e.GameObjects()[charGameObj.Properties()["claim_obelisk_id"].(string)]
		if obelisk == nil {
			e.SendSystemMessage("You don't have a claim.", player)
			return false
		}
		Move(e, charGameObj, obelisk.X() + 1.0, obelisk.Y() + 1.0)
		e.SendGameObjectUpdate(charGameObj, "update_object")
	}

	return true
}