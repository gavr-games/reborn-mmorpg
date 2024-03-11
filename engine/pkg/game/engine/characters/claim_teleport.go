package characters

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// This func is called via delayed action mechanism
// params: playerId
func ClaimTeleport(e entity.IEngine, params map[string]interface{}) bool {
	playerId := int(params["playerId"].(float64))
	if player, ok := e.Players().Load(playerId); ok {
		charGameObj := e.GameObjects()[player.CharacterGameObjectId]
		charGameObj.(entity.ICharacterObject).DeselectTarget(e)
		obelisk := e.GameObjects()[charGameObj.Properties()["claim_obelisk_id"].(string)]
		if obelisk == nil {
			e.SendSystemMessage("You don't have a claim.", player)
			return false
		}
		charGameObj.(entity.ICharacterObject).Move(e, obelisk.X() + 1.0, obelisk.Y() + 1.0)
		e.SendGameObjectUpdate(charGameObj, "update_object")
	}

	return true
}