package plants

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// This func is called via delayed action mechanism
// params: game_object_id
func Grow(e entity.IEngine, params map[string]interface{}) bool {
	if plant, plantOk := e.GameObjects().Load(params["game_object_id"].(string)); plantOk {
		return plant.(entity.IPlantObject).Grow(e)
	} else {
		return false
	}
}
