package claims

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// This func is called via delayed action mechanism
// params: game_object_id
func Init(e entity.IEngine, params map[string]interface{}) bool {
	claimObeliskId := params["game_object_id"].(string)
	if obelisk, ok := e.GameObjects()[claimObeliskId]; ok {
		return obelisk.(entity.IClaimObeliskObject).Init(e)
	} else {
		return false
	}
}