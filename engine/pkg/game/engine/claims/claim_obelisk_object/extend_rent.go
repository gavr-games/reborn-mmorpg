package claim_obelisk_object

import (
	"fmt"

	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/constants"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/claims"
)

func (claimObelisk *ClaimObeliskObject) ExtendRent(e entity.IEngine) bool {
	charGameObj := e.GameObjects()[claimObelisk.Properties()["crafted_by_character_id"].(string)]
	if charGameObj == nil {
		return false
	}

	slots := charGameObj.Properties()["slots"].(map[string]interface{})

	playerId := charGameObj.Properties()["player_id"].(int)
	player := e.Players()[playerId]
	if player == nil {
		return false
	}

	// Check claim access
	if !claims.CheckAccess(e, charGameObj, claimObelisk) {
		e.SendSystemMessage("You don't have an access to this claim.", player)
		return false
	}

	// check container has money
	container := e.GameObjects()[slots["back"].(string)]
	if !container.(entity.IContainerObject).HasItemsKinds(e, map[string]interface{}{
		"gold": constants.ClaimRentCost,
	}) {
		e.SendSystemMessage(fmt.Sprintf("You need %d gold to pay rent.", int(constants.ClaimRentCost)), player)
		return false
	}

	// substract money
	if !container.(entity.IContainerObject).RemoveItemsKinds(e, player, map[string]interface{}{
		"gold": constants.ClaimRentCost,
	}) {
		e.SendSystemMessage("Can't remove required gold.", player)
		return false
	}

	// replace delayed action
	claimObelisk.Properties()["payed_until"] = claimObelisk.Properties()["payed_until"].(float64) + constants.ClaimRentDuration
	delayedAction := &entity.DelayedAction{
		FuncName: "ExpireClaim",
		Params: map[string]interface{}{
			"claim_obelisk_id": claimObelisk.Id(),
		},
		TimeLeft: claimObelisk.Properties()["payed_until"].(float64) - float64(utils.MakeTimestamp()),
	}
	claimObelisk.SetCurrentAction(delayedAction)

	storage.GetClient().Updates <- claimObelisk.Clone()
	e.SendSystemMessage("You have payed the rent.", player)

	return true
}