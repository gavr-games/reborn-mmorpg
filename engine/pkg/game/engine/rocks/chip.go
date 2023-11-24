package rocks

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// This func is called via delayed action mechanism
// params: characterId, rockId
func Chip(e entity.IEngine, params map[string]interface{}) bool {
	rock := e.GameObjects()[params["rockId"].(string)].(entity.IRockObject)
	character := e.GameObjects()[params["characterId"].(string)]
	return rock.Chip(e, character)
}
