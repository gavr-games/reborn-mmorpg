package claim_obelisk_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/claims"
)

func (claimObelisk *ClaimObeliskObject) Destroy(e entity.IEngine, player *entity.Player) bool {
	charGameObj := e.GameObjects()[player.CharacterGameObjectId]

	// Check claim access
	if !claims.CheckAccess(e, charGameObj, claimObelisk) {
		e.SendSystemMessage("You don't have an access to this claim.", player)
		return false
	}

	// Check near building
	if !claimObelisk.IsCloseTo(charGameObj) {
		e.SendSystemMessage("You need to be closer to the item.", player)
		return false
	}

	// Destroy
	claimObelisk.Remove(e)

	return true
}
