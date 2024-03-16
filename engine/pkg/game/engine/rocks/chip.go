package rocks

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// This func is called via delayed action mechanism
// params: characterId, rockId
func Chip(e entity.IEngine, params map[string]interface{}) bool {
	var (
		rock, character entity.IGameObject
		rockOk, charOk bool
	)
	if rock, rockOk = e.GameObjects().Load(params["rockId"].(string)); !rockOk {
		return false
	}
	if character, charOk = e.GameObjects().Load(params["characterId"].(string)); !charOk {
		return false
	}
	return rock.(entity.IRockObject).Chip(e, character)
}
