package claims

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/constants"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

// This func is called via delayed action mechanism
// params: game_object_id
func Init(e entity.IEngine, params map[string]interface{}) bool {
	claimObeliskId := params["game_object_id"].(string)
	if obelisk, ok := e.GameObjects()[claimObeliskId]; ok {
		charGameObj := e.GameObjects()[obelisk.Properties()["crafted_by_character_id"].(string)]
		if charGameObj == nil {
			return false
		}

		// Create claim area
		additionalProps := make(map[string]interface{})
		additionalProps["claim_obelisk_id"] = obelisk.Id
		claimArea := e.CreateGameObject("claim/claim_area", obelisk.X() - constants.ClaimArea / 2, obelisk.Y() - constants.ClaimArea / 2, 0.0, obelisk.Floor(), additionalProps)
		
		obelisk.Properties()["claim_area_id"] = claimArea.Id

		// Init rent
		obelisk.Properties()["payed_until"] = float64(utils.MakeTimestamp()) + constants.ClaimRentDuration

		delayedAction := &entity.DelayedAction{
			FuncName: "ExpireClaim",
			Params: map[string]interface{}{
				"claim_obelisk_id": obelisk.Id(),
			},
			TimeLeft: constants.ClaimRentDuration,
		}
		obelisk.SetCurrentAction(delayedAction)

		// Set claim obelisk id for character
		charGameObj.Properties()["claim_obelisk_id"] = obelisk.Id()

		storage.GetClient().Updates <- obelisk.Clone()
		storage.GetClient().Updates <- charGameObj.Clone()

		e.SendGameObjectUpdate(claimArea, "add_object")
	} else {
		return false
	}
	return true
}