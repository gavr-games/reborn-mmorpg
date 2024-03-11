package hatcheries

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// This func is called via delayed action mechanism
// params: hatcheryId, mobPath
func Hatch(e entity.IEngine, params map[string]interface{}) bool {
	if hatchery, ok := e.GameObjects().Load(params["hatcheryId"].(string)); ok {
		mobPath := params["mobPath"].(string)
		return hatchery.(entity.IHatcheryObject).Hatch(e, mobPath)
	} else {
		return false
	}
}
