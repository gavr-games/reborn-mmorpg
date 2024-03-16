package claim_obelisk_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

func (claimObelisk *ClaimObeliskObject) Remove(e entity.IEngine) bool {
	var (
		charGameObj, claimAreaObj entity.IGameObject
		charOk, areaOk bool
	)
	if charGameObj, charOk = e.GameObjects().Load(claimObelisk.GetProperty("crafted_by_character_id").(string)); !charOk {
		return false
	}

	// remove claim area
	if claimAreaObj, areaOk = e.GameObjects().Load(claimObelisk.GetProperty("claim_area_id").(string)); !areaOk {
		return false
	}
	e.Floors()[claimAreaObj.Floor()].FilteredRemove(claimAreaObj, func(b utils.IBounds) bool {
		return claimAreaObj.Id() == b.(entity.IGameObject).Id()
	})
	e.GameObjects().Delete(claimAreaObj.Id())
	e.SendGameObjectUpdate(claimAreaObj, "remove_object")

	// remove obelisk
	e.Floors()[claimObelisk.Floor()].FilteredRemove(claimObelisk, func(b utils.IBounds) bool {
		return claimObelisk.Id() == b.(entity.IGameObject).Id()
	})
	e.GameObjects().Delete(claimObelisk.Id())
	e.SendGameObjectUpdate(claimObelisk, "remove_object")

	// remove obelisk from character
	charGameObj.SetProperty("claim_obelisk_id", nil)
	storage.GetClient().Updates <- charGameObj.Clone()

	return true
}
