package claims

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// This func is called via delayed action mechanism
// params: claim_obelisk_id
func Expire(e entity.IEngine, params map[string]interface{}) bool {
	claimObeliskId := params["claim_obelisk_id"].(string)
	if obelisk, ok := e.GameObjects()[claimObeliskId]; ok {
		return obelisk.(entity.IClaimObeliskObject).Remove(e)
	} else {
		return false
	}
}