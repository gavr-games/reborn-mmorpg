package plant_object

import (
	"fmt"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

func (plant *PlantObject) Cut(e entity.IEngine, charGameObj entity.IGameObject) bool {
	playerId := charGameObj.Properties()["player_id"].(int)
	if player, ok := e.Players()[playerId]; ok {
		slots := charGameObj.Properties()["slots"].(map[string]interface{})

		if slots["back"] == nil {
			e.SendSystemMessage("You don't have container", player)
			return false
		}

		if !plant.CheckCut(e, charGameObj) {
			return false
		}

		// Check near the plant
		if !plant.IsCloseTo(charGameObj) {
			e.SendSystemMessage("You need to be closer to the plant.", player)
			return false
		}

		// Create plant resource
		resources := plant.Properties()["resources"].(map[string]interface{})
		resource_key := ""
		for k, _ := range resources {
			resource_key = k
			break
		}
		resourceObj := e.CreateGameObject(fmt.Sprintf("resource/%s", resource_key), charGameObj.X(), charGameObj.Y(), 0.0, -1, nil)

		// put resource to container or drop it to the ground
		container := e.GameObjects()[slots["back"].(string)]
		container.(entity.IContainerObject).PutOrDrop(e, charGameObj, resourceObj.Id(), -1)

		// Decrease resources stored in the cactus
		resources[resource_key] = resources[resource_key].(float64) - 1.0

		// Remove plant if no resource inside
		if resources[resource_key].(float64) <= 0 {
			e.SendGameObjectUpdate(plant, "remove_object")

			e.Floors()[plant.Floor()].FilteredRemove(e.GameObjects()[plant.Id()], func(b utils.IBounds) bool {
				return plant.Id() == b.(entity.IGameObject).Id()
			})
			e.GameObjects()[plant.Id()] = nil
			delete(e.GameObjects(), plant.Id())
		} else {
			storage.GetClient().Updates <- plant.Clone()
		}

		e.SendSystemMessage(fmt.Sprintf("You received a %s.", resourceObj.Kind()), player)
	} else {
		return false
	}

	return true
}
