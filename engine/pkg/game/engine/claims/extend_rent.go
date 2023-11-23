package claims

import (
	"fmt"

	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/constants"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/containers"
)

func ExtendRent(e entity.IEngine, claimObeliskId string) bool {
	if obelisk, ok := e.GameObjects()[claimObeliskId]; ok {
		charGameObj := e.GameObjects()[obelisk.Properties()["crafted_by_character_id"].(string)]
		if charGameObj == nil {
			return false
		}

		slots := charGameObj.Properties()["slots"].(map[string]interface{})

		playerId := charGameObj.Properties()["player_id"].(int)
		player := e.Players()[playerId]
		if player == nil {
			return false
		}

		// check is obelisk
		if obelisk.Properties()["kind"].(string) != "claim_obelisk" {
			e.SendSystemMessage("Please choose claim obelisk.", player)
			return false
		}

		// Check claim access
		if !CheckAccess(e, charGameObj, obelisk) {
			e.SendSystemMessage("You don't have an access to this claim.", player)
			return false
		}

		// check container has money
		if !containers.HasItemsKinds(e, slots["back"].(string), map[string]interface{}{
			"gold": constants.ClaimRentCost,
		}) {
			e.SendSystemMessage(fmt.Sprintf("You need %d gold to pay rent.", int(constants.ClaimRentCost)), player)
			return false
		}

		// substract money
		if !containers.RemoveItemsKinds(e, player, slots["back"].(string), map[string]interface{}{
			"gold": constants.ClaimRentCost,
		}) {
			e.SendSystemMessage("Can't remove required gold.", player)
			return false
		}

		// replace delayed action
		obelisk.Properties()["payed_until"] = obelisk.Properties()["payed_until"].(float64) + constants.ClaimRentDuration
		delayedAction := &entity.DelayedAction{
			FuncName: "ExpireClaim",
			Params: map[string]interface{}{
				"claim_obelisk_id": obelisk.Id(),
			},
			TimeLeft: obelisk.Properties()["payed_until"].(float64) - float64(utils.MakeTimestamp()),
		}
		obelisk.SetCurrentAction(delayedAction)

		storage.GetClient().Updates <- obelisk.Clone()
		e.SendSystemMessage("You have payed the rent.", player)
	} else {
		return false
	}
	return true
}