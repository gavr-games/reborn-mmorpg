package rock_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

func (rock *RockObject) Chip(e entity.IEngine, charGameObj entity.IGameObject) bool {
	playerId := charGameObj.Properties()["player_id"].(int)
	if player, ok := e.Players()[playerId]; ok {
		slots := charGameObj.Properties()["slots"].(map[string]interface{})

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
		stoneObj := e.CreateGameObject("resource/stone", charGameObj.X(), charGameObj.Y(), 0.0, -1, nil)

		container := e.GameObjects()[slots["back"].(string)]
		container.(entity.IContainerObject).PutOrDrop(e, charGameObj, stoneObj.Id(), -1)

		// Decrease stones stored in the rock
		resources := rock.Properties()["resources"].(map[string]interface{})
		resources["stone"] = resources["stone"].(float64) - 1.0

		// Remove rock if no stones inside
		if resources["stone"].(float64) <= 0 {
			e.SendGameObjectUpdate(rock, "remove_object")

			e.Floors()[rock.Floor()].FilteredRemove(e.GameObjects()[rock.Id()], func(b utils.IBounds) bool {
				return rock.Id() == b.(entity.IGameObject).Id()
			})
			e.GameObjects()[rock.Id()] = nil
			delete(e.GameObjects(), rock.Id())
		} else {
			storage.GetClient().Updates <- rock.Clone()
		}

		e.SendSystemMessage("You received a stone.", player)
	} else {
		return false
	}

	return true
}
