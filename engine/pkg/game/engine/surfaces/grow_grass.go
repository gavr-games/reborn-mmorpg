package surfaces

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// This func is called via delayed action mechanism
// params: game_object_id
func GrowGrass(e entity.IEngine, params map[string]interface{}) bool {
	if dirt, dirtOk := e.GameObjects().Load(params["game_object_id"].(string)); dirtOk {
		// Add grass
		e.CreateGameObject("surface/grass", dirt.X(), dirt.Y(), 0.0, dirt.GameAreaId(), nil)
		/*
		// Frontend shows the default grass plane, so for performance optimization we can skip this
		grass := e.CreateGameObject("surface/grass", dirt.X(), dirt.Y(), 0.0, dirt.GameAreaId(), nil)
		e.SendResponseToVisionAreas(grass, "add_object", map[string]interface{}{
			"object": grass.Clone(),
		})
		*/

		// Remove dirt
		e.RemoveGameObject(dirt)

		return true
	} else {
		return false
	}
}