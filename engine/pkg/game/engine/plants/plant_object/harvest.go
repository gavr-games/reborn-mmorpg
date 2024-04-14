package plant_object

import (
	"fmt"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (plant *PlantObject) Harvest(e entity.IEngine, charGameObj entity.IGameObject) bool {
	playerId := charGameObj.GetProperty("player_id").(int)
	if player, ok := e.Players().Load(playerId); ok {
		slots := charGameObj.GetProperty("slots").(map[string]interface{})

		if slots["back"] == nil {
			e.SendSystemMessage("You don't have container", player)
			return false
		}

		if !plant.CheckHarvest(e, charGameObj) {
			return false
		}

		// Check near the plant
		if !plant.IsCloseTo(charGameObj) {
			e.SendSystemMessage("You need to be closer to the plant.", player)
			return false
		}

		// Create harvested resources
		var (
			container entity.IGameObject
			contOk bool
		)
		if container, contOk = e.GameObjects().Load(slots["back"].(string)); !contOk {
			return false
		}
		resources := plant.GetProperty("resources").(map[string]interface{})
		for resourceKey, amount := range resources {
			for i := 0; i < int(amount.(float64)); i++ {
				resourceObj := e.CreateGameObject(fmt.Sprintf("resource/%s", resourceKey), charGameObj.X(), charGameObj.Y(), 0.0, "", nil)
				// put resource to container or drop it to the ground
				container.(entity.IContainerObject).PutOrDrop(e, charGameObj, resourceObj.Id(), -1)
				e.SendSystemMessage(fmt.Sprintf("You received a %s.", resourceObj.Kind()), player)
			}
		}

		// Remove plant
		e.RemoveGameObject(plant)

		charGameObj.(entity.ILevelingObject).AddExperience(e, "harvest_plant")

		e.SendSystemMessage(fmt.Sprintf("You harvested a %s.", plant.Kind()), player)
	} else {
		return false
	}

	return true
}