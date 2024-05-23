package claim_obelisk_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/constants"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (claimObelisk *ClaimObeliskObject) Init(e entity.IEngine) bool {
	var (
		charGameObj entity.IGameObject
		charOk bool
	)
	if charGameObj, charOk = e.GameObjects().Load(claimObelisk.GetProperty("crafted_by_character_id").(string)); !charOk {
		return false
	}

	// Create claim area
	additionalProps := make(map[string]interface{})
	additionalProps["claim_obelisk_id"] = claimObelisk.Id()
	claimArea := e.CreateGameObject("claim/claim_area", claimObelisk.X() - constants.ClaimArea / 2, claimObelisk.Y() - constants.ClaimArea / 2, 0.0, claimObelisk.GameAreaId(), additionalProps)

	claimObelisk.SetProperty("claim_area_id", claimArea.Id())

	// Init rent
	claimObelisk.SetProperty("payed_until", float64(utils.MakeTimestamp()) + constants.ClaimRentDuration)

	delayedAction := entity.NewDelayedAction(
		"ExpireClaim",
		map[string]interface{}{
			"claim_obelisk_id": claimObelisk.Id(),
		},
		constants.ClaimRentDuration,
		entity.DelayedActionReady,
	)
	claimObelisk.SetCurrentAction(delayedAction)
	e.DelayedActions().Store(claimObelisk.Id(), claimObelisk)

	// Set claim obelisk id for character
	charGameObj.SetProperty("claim_obelisk_id", claimObelisk.Id())

	storage.GetClient().Updates <- claimObelisk.Clone()
	storage.GetClient().Updates <- charGameObj.Clone()

	e.SendGameObjectUpdate(claimArea, "add_object")

	return true
}
