package claim_obelisk_object

import (
	"fmt"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/constants"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/claims"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

func (claimObelisk *ClaimObeliskObject) ExtendRent(e entity.IEngine) bool {
	var (
		charGameObj, container entity.IGameObject
		charOk, contOk bool
	)
	if charGameObj, charOk = e.GameObjects().Load(claimObelisk.GetProperty("crafted_by_character_id").(string)); !charOk {
		return false
	}

	slots := charGameObj.GetProperty("slots").(map[string]interface{})

	playerId := charGameObj.GetProperty("player_id").(int)
	player, ok := e.Players().Load(playerId)
	if player == nil || !ok {
		return false
	}

	// Check claim access
	if !claims.CheckAccess(e, charGameObj, claimObelisk) {
		e.SendSystemMessage("You don't have an access to this claim.", player)
		return false
	}

	// check container has money
	if container, contOk = e.GameObjects().Load(slots["back"].(string)); !contOk {
		return false
	}
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
	claimObelisk.SetProperty("payed_until", claimObelisk.GetProperty("payed_until").(float64) + constants.ClaimRentDuration)
	delayedAction := entity.NewDelayedAction(
		"ExpireClaim",
		map[string]interface{}{
			"claim_obelisk_id": claimObelisk.Id(),
		},
		claimObelisk.GetProperty("payed_until").(float64) - float64(utils.MakeTimestamp()),
		entity.DelayedActionReady,
	)
	claimObelisk.SetCurrentAction(delayedAction)

	storage.GetClient().Updates <- claimObelisk.Clone()
	e.SendSystemMessage("You have payed the rent.", player)

	return true
}
