package shovels

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// This func is called via delayed action mechanism
// params: characterId, shovelId
func Dig(e entity.IEngine, params map[string]interface{}) bool {
	var (
		shovel, character entity.IGameObject
		shovelOk, charOk bool
	)
	if shovel, shovelOk = e.GameObjects().Load(params["shovelId"].(string)); !shovelOk {
		return false
	}
	if character, charOk = e.GameObjects().Load(params["characterId"].(string)); !charOk {
		return false
	}
	return shovel.(entity.IShovelObject).Dig(e, character)
}