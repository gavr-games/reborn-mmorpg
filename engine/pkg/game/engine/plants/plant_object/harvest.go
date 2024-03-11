package plant_object

import (
	"fmt"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

func (plant *PlantObject) Harvest(e entity.IEngine, charGameObj entity.IGameObject) bool {
	playerId := charGameObj.Properties()["player_id"].(int)
	if player, ok := e.Players().Load(playerId); ok {
		slots := charGameObj.Properties()["slots"].(map[string]interface{})

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
		resources := plant.Properties()["resources"].(map[string]interface{})
		for resourceKey, amount := range resources {
			for i := 0; i < int(amount.(float64)); i++ {
				resourceObj := e.CreateGameObject(fmt.Sprintf("resource/%s", resourceKey), charGameObj.X(), charGameObj.Y(), 0.0, -1, nil)
				// put resource to container or drop it to the ground
				container.(entity.IContainerObject).PutOrDrop(e, charGameObj, resourceObj.Id(), -1)
				e.SendSystemMessage(fmt.Sprintf("You received a %s.", resourceObj.Kind()), player)
			}
		}

		// Remove plant
		e.SendGameObjectUpdate(plant, "remove_object")
		e.Floors()[plant.Floor()].FilteredRemove(plant, func(b utils.IBounds) bool {
			return plant.Id() == b.(entity.IGameObject).Id()
		})
		e.GameObjects().Delete(plant.Id())

		e.SendSystemMessage(fmt.Sprintf("You harvested a %s.", plant.Kind()), player)
	} else {
		return false
	}

	return true
}