package claim_obelisk_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/claims"
)

func (claimObelisk *ClaimObeliskObject) Destroy(e entity.IEngine, player *entity.Player) bool {
	var (
		charGameObj entity.IGameObject
		charOk bool
	)
	if charGameObj, charOk = e.GameObjects().Load(player.CharacterGameObjectId); !charOk {
		return false
	}

	// Check claim access
	if !claims.CheckAccess(e, charGameObj, claimObelisk) {
		e.SendSystemMessage("You don't have an access to this claim.", player)
		return false
	}

	// Check near building
	if !claimObelisk.IsCloseTo(charGameObj) {
		e.SendSystemMessage("You need to be closer to the claim.", player)
		return false
	}

	// Destroy
	claimObelisk.Remove(e)

	return true
}
