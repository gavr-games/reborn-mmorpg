package cactus_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

func (cactus *CactusObject) Cut(e entity.IEngine, charGameObj entity.IGameObject) bool {
	playerId := charGameObj.Properties()["player_id"].(int)
	if player, ok := e.Players()[playerId]; ok {
		slots := charGameObj.Properties()["slots"].(map[string]interface{})

		if slots["back"] == nil {
			e.SendSystemMessage("You don't have container", player)
			return false
		}

		if !cactus.CheckCut(e, charGameObj) {
			return false
		}

		// Check near the cactus
		if !cactus.IsCloseTo(charGameObj) {
			e.SendSystemMessage("You need to be closer to the cactus.", player)
			return false
		}

		// Create cactus slice
		sliceObj := e.CreateGameObject("resource/cactus_slice", charGameObj.X(), charGameObj.Y(), 0.0, -1, nil)

		// put log to container or drop it to the ground
		container := e.GameObjects()[slots["back"].(string)]
		container.(entity.IContainerObject).PutOrDrop(e, charGameObj, sliceObj.Id(), -1)

		// Decrease slices stored in the cactus
		resources := cactus.Properties()["resources"].(map[string]interface{})
		resources["cactus_slice"] = resources["cactus_slice"].(float64) - 1.0

		// Remove cactus if no cactus_slice inside
		if resources["cactus_slice"].(float64) <= 0 {
			e.SendGameObjectUpdate(cactus, "remove_object")

			e.Floors()[cactus.Floor()].FilteredRemove(e.GameObjects()[cactus.Id()], func(b utils.IBounds) bool {
				return cactus.Id() == b.(entity.IGameObject).Id()
			})
			e.GameObjects()[cactus.Id()] = nil
			delete(e.GameObjects(), cactus.Id())
		} else {
			storage.GetClient().Updates <- cactus.Clone()
		}

		e.SendSystemMessage("You received a cactus slice.", player)
	} else {
		return false
	}

	return true
}
