package claims

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

// This func is called via delayed action mechanism
// params: claim_obelisk_id
func Expire(e entity.IEngine, params map[string]interface{}) bool {
	claimObeliskId := params["claim_obelisk_id"].(string)
	if obelisk, ok := e.GameObjects()[claimObeliskId]; ok {
		charGameObj := e.GameObjects()[obelisk.Properties["crafted_by_character_id"].(string)]
		if charGameObj == nil {
			return false
		}

		playerId := charGameObj.Properties["player_id"].(int)
		player := e.Players()[playerId]
		if player == nil {
			return false
		}

		// remove claim area
		claimAreaObj := e.GameObjects()[obelisk.Properties["claim_area_id"].(string)]
		e.Floors()[claimAreaObj.Floor].FilteredRemove(claimAreaObj, func(b utils.IBounds) bool {
			return claimAreaObj.Id == b.(*entity.GameObject).Id
		})
		e.GameObjects()[claimAreaObj.Id] = nil
		delete(e.GameObjects(), claimAreaObj.Id)
		storage.GetClient().Deletes <- claimAreaObj.Id
		e.SendGameObjectUpdate(claimAreaObj, "remove_object")

		// remove obelisk
		e.Floors()[obelisk.Floor].FilteredRemove(obelisk, func(b utils.IBounds) bool {
			return claimObeliskId == b.(*entity.GameObject).Id
		})
		e.GameObjects()[claimObeliskId] = nil
		delete(e.GameObjects(), claimObeliskId)
		storage.GetClient().Deletes <- claimObeliskId
		e.SendGameObjectUpdate(obelisk, "remove_object")

		// remove obelisk from character
		charGameObj.Properties["claim_obelisk_id"] = nil
		storage.GetClient().Updates <- charGameObj.Clone()
	} else {
		return false
	}
	return true
}