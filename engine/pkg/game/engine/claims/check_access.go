package claims

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// Check that charGameObj can interract with targetObj
// it can interract if the targetObj is not inside the claim area of another player
func CheckAccess(e entity.IEngine, charGameObj entity.IGameObject, targetObj entity.IGameObject) bool {
	// Check not intersecting with claim areas
	if obelisk := GetClaimObelisk(e, targetObj); obelisk != nil {
		if owner, ownerOk := e.GameObjects().Load(obelisk.GetProperty("crafted_by_character_id").(string)); ownerOk {
			return owner.Id() == charGameObj.Id()
		} else {
			return false
		}
	}

	return true
}
