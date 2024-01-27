package surfaces

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// This func is called via delayed action mechanism
// params: game_object_id
func GrowGrass(e entity.IEngine, params map[string]interface{}) bool {
	dirt := e.GameObjects()[params["game_object_id"].(string)]

	// Add grass
	grass := e.CreateGameObject("surface/grass", dirt.X(), dirt.Y(), 0.0, dirt.Floor(), nil)
	e.SendResponseToVisionAreas(grass, "add_object", map[string]interface{}{
		"object": grass,
	})

	// Remove dirt
	e.SendGameObjectUpdate(dirt, "remove_object")
	e.Floors()[dirt.Floor()].FilteredRemove(dirt, func(b utils.IBounds) bool {
		return dirt.Id() == b.(entity.IGameObject).Id()
	})
	e.GameObjects()[dirt.Id()] = nil
	delete(e.GameObjects(), dirt.Id())
	
	return true
}