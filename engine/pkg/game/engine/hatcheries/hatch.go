package hatcheries

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// This func is called via delayed action mechanism
// params: hatcheryId, mobPath
func Hatch(e entity.IEngine, params map[string]interface{}) bool {
	hatchery := e.GameObjects()[params["hatcheryId"].(string)].(entity.IHatcheryObject)
	mobPath := params["mobPath"].(string)
	return hatchery.Hatch(e, mobPath)
}
