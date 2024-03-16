package plants

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// This func is called via delayed action mechanism
// params: characterId, plantId
func Harvest(e entity.IEngine, params map[string]interface{}) bool {
	var (
		plant, character entity.IGameObject
		plantOk, charOk bool
	)
	if plant, plantOk = e.GameObjects().Load(params["plantId"].(string)); !plantOk {
		return false
	}
	if character, charOk = e.GameObjects().Load(params["characterId"].(string)); !charOk {
		return false
	}
	return plant.(entity.IPlantObject).Harvest(e, character)
}
