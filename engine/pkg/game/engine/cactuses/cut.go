package cactuses

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// This func is called via delayed action mechanism
// params: characterId, cactusId
func Cut(e entity.IEngine, params map[string]interface{}) bool {
	cactus := e.GameObjects()[params["cactusId"].(string)].(entity.ICactusObject)
	character := e.GameObjects()[params["characterId"].(string)]
	return cactus.Cut(e, character)
}
