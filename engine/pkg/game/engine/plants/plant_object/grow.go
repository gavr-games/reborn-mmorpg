package plant_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (plant *PlantObject) Grow(e entity.IEngine) bool {
	// Create next plant
	if growsInto := plant.GetProperty("grows_into"); growsInto != nil {
		nextPlant := e.CreateGameObject(growsInto.(string), plant.X(), plant.Y(), 0.0, plant.GameAreaId(), nil)
		e.SendResponseToVisionAreas(nextPlant, "add_object", map[string]interface{}{
			"object": nextPlant.Clone(),
		})
	}

	// Remove plant
	e.RemoveGameObject(plant)

	return true
}
