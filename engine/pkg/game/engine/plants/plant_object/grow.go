package plant_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (plant *PlantObject) Grow(e entity.IEngine) bool {
	// Create next plant
	if growsInto, ok := plant.Properties()["grows_into"]; ok {
		nextPlant := e.CreateGameObject(growsInto.(string), plant.X(), plant.Y(), 0.0, plant.Floor(), nil)
		e.SendResponseToVisionAreas(nextPlant, "add_object", map[string]interface{}{
			"object": nextPlant,
		})
	}

	// Remove plant
	e.SendGameObjectUpdate(plant, "remove_object")

	e.Floors()[plant.Floor()].FilteredRemove(e.GameObjects()[plant.Id()], func(b utils.IBounds) bool {
		return plant.Id() == b.(entity.IGameObject).Id()
	})
	e.GameObjects()[plant.Id()] = nil
	delete(e.GameObjects(), plant.Id())

	return true
}
