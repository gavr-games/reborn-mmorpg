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

		if cactus == nil {
			e.SendSystemMessage("Cactus does not exist.", player)
			return false
		}

		// Create cactus slice
		sliceObj := e.CreateGameObject("resource/cactus_slice", charGameObj.X(), charGameObj.Y(), 0.0, -1, nil)

		// check character has container
		putInContainer := false
		if (slots["back"] != nil) {
			// put log to container
			container := e.GameObjects()[slots["back"].(string)]
			putInContainer = container.(entity.IContainerObject).Put(e, player, sliceObj.Id(), -1)
		}

		// OR drop logs on the ground
		if !putInContainer {
			sliceObj.SetFloor(charGameObj.Floor())
			e.Floors()[sliceObj.Floor()].Insert(sliceObj)
			storage.GetClient().Updates <- sliceObj.Clone()
			e.SendResponseToVisionAreas(e.GameObjects()[player.CharacterGameObjectId], "add_object", map[string]interface{}{
				"object": sliceObj,
			})
		}

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
