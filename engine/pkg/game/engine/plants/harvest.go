package plants

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// This func is called via delayed action mechanism
// params: characterId, plantId
func Harvest(e entity.IEngine, params map[string]interface{}) bool {
	plant := e.GameObjects()[params["plantId"].(string)].(entity.IPlantObject)
	character := e.GameObjects()[params["characterId"].(string)]
	return plant.Harvest(e, character)
}
