package plants

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// This func is called via delayed action mechanism
// params: game_object_id
func Grow(e entity.IEngine, params map[string]interface{}) bool {
	plant := e.GameObjects()[params["game_object_id"].(string)].(entity.IPlantObject)
	return plant.Grow(e)
}
