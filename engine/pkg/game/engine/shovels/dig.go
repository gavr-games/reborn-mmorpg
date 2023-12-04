package shovels

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// This func is called via delayed action mechanism
// params: characterId, shovelId
func Dig(e entity.IEngine, params map[string]interface{}) bool {
	shovel := e.GameObjects()[params["shovelId"].(string)].(entity.IShovelObject)
	character := e.GameObjects()[params["characterId"].(string)]
	return shovel.Dig(e, character)
}