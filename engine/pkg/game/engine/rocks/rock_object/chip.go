package rock_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

func (rock *RockObject) Chip(e entity.IEngine, charGameObj entity.IGameObject) bool {
	playerId := charGameObj.GetProperty("player_id").(int)
	if player, ok := e.Players().Load(playerId); ok {
		slots := charGameObj.GetProperty("slots").(map[string]interface{})

		if slots["back"] == nil {
			e.SendSystemMessage("You don't have container", player)
			return false
		}

		if !rock.CheckChip(e, charGameObj) {
			return false
		}

		// Check near the rock
		if !rock.IsCloseTo(charGameObj) {
			e.SendSystemMessage("You need to be closer to the rock.", player)
			return false
		}

		// Create stone
		stoneObj := e.CreateGameObject("resource/stone", charGameObj.X(), charGameObj.Y(), 0.0, "", nil)

		if container, contOk := e.GameObjects().Load(slots["back"].(string)); contOk {
			container.(entity.IContainerObject).PutOrDrop(e, charGameObj, stoneObj.Id(), -1)
		} else {
			return false
		}

		// Decrease stones stored in the rock
		resources := rock.GetProperty("resources").(map[string]interface{})
		resources["stone"] = resources["stone"].(float64) - 1.0
		rock.SetProperty("resources", resources)

		// Remove rock if no stones inside
		if resources["stone"].(float64) <= 0 {
			e.RemoveGameObject(rock)
		} else {
			storage.GetClient().Updates <- rock.Clone()
		}

		charGameObj.(entity.ILevelingObject).AddExperience(e, "chip_rock")

		e.SendSystemMessage("You received a stone.", player)
	} else {
		return false
	}

	return true
}
