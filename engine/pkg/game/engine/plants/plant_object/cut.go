package plant_object

import (
	"fmt"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

func (plant *PlantObject) Cut(e entity.IEngine, charGameObj entity.IGameObject) bool {
	playerId := charGameObj.GetProperty("player_id").(int)
	if player, ok := e.Players().Load(playerId); ok {
		slots := charGameObj.GetProperty("slots").(map[string]interface{})

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
		resources := plant.GetProperty("resources").(map[string]interface{})
		resourceKey := ""
		for k, _ := range resources {
			resourceKey = k
			break
		}
		resourceObj := e.CreateGameObject(fmt.Sprintf("resource/%s", resourceKey), charGameObj.X(), charGameObj.Y(), 0.0, -1, nil)

		// put resource to container or drop it to the ground
		var (
			container entity.IGameObject
			contOk bool
		)
		if container, contOk = e.GameObjects().Load(slots["back"].(string)); !contOk {
			return false
		}
		container.(entity.IContainerObject).PutOrDrop(e, charGameObj, resourceObj.Id(), -1)

		// Decrease resources stored in the cactus
		resources[resourceKey] = resources[resourceKey].(float64) - 1.0
		plant.SetProperty("resources", resources)

		// Remove plant if no resource inside
		if resources[resourceKey].(float64) <= 0 {
			e.SendGameObjectUpdate(plant, "remove_object")

			e.Floors()[plant.Floor()].FilteredRemove(plant, func(b utils.IBounds) bool {
				return plant.Id() == b.(entity.IGameObject).Id()
			})
			e.GameObjects().Delete(plant.Id())
		} else {
			storage.GetClient().Updates <- plant.Clone()
		}

		e.SendSystemMessage(fmt.Sprintf("You received a %s.", resourceObj.Kind()), player)
	} else {
		return false
	}

	return true
}
