package claim_obelisk_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

func (claimObelisk *ClaimObeliskObject) Remove(e entity.IEngine) bool {
	charGameObj := e.GameObjects()[claimObelisk.Properties()["crafted_by_character_id"].(string)]
	if charGameObj == nil {
		return false
	}

	// remove claim area
	claimAreaObj := e.GameObjects()[claimObelisk.Properties()["claim_area_id"].(string)]
	e.Floors()[claimAreaObj.Floor()].FilteredRemove(claimAreaObj, func(b utils.IBounds) bool {
		return claimAreaObj.Id() == b.(entity.IGameObject).Id()
	})
	e.GameObjects()[claimAreaObj.Id()] = nil
	delete(e.GameObjects(), claimAreaObj.Id())
	e.SendGameObjectUpdate(claimAreaObj, "remove_object")

	// remove obelisk
	e.Floors()[claimObelisk.Floor()].FilteredRemove(claimObelisk, func(b utils.IBounds) bool {
		return claimObelisk.Id() == b.(entity.IGameObject).Id()
	})
	e.GameObjects()[claimObelisk.Id()] = nil
	delete(e.GameObjects(), claimObelisk.Id())
	e.SendGameObjectUpdate(claimObelisk, "remove_object")

	// remove obelisk from character
	charGameObj.Properties()["claim_obelisk_id"] = nil
	storage.GetClient().Updates <- charGameObj.Clone()

	return true
}
