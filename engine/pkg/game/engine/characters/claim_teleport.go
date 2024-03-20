package characters

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// This func is called via delayed action mechanism
// params: playerId
func ClaimTeleport(e entity.IEngine, params map[string]interface{}) bool {
	playerId := int(params["playerId"].(float64))
	if player, ok := e.Players().Load(playerId); ok {
		if charGameObj, charOk := e.GameObjects().Load(player.CharacterGameObjectId); charOk {
			charGameObj.(entity.ICharacterObject).DeselectTarget(e)
			if obelisk, obeliskOk := e.GameObjects().Load(charGameObj.GetProperty("claim_obelisk_id").(string)); obeliskOk {
				if obelisk == nil {
					e.SendSystemMessage("You don't have a claim.", player)
					return false
				}
				// Send remove to all players who see character
				charGameObjClone := charGameObj.Clone()
				e.SendResponseToVisionAreas(charGameObjClone, "remove_object", map[string]interface{}{
					"object": charGameObjClone,
				})
				charGameObj.(entity.ICharacterObject).Move(e, obelisk.X() + 1.0, obelisk.Y() + 1.0, obelisk.Floor())
			}
			e.SendGameObjectUpdate(charGameObj, "update_object")
		}
	}

	return true
}