package rock_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/containers"
)

func (rock *RockObject) Chip(e entity.IEngine, charGameObj entity.IGameObject) bool {
	playerId := charGameObj.Properties()["player_id"].(int)
	if player, ok := e.Players()[playerId]; ok {
		slots := charGameObj.Properties()["slots"].(map[string]interface{})

		// Create log
		stoneObj := e.CreateGameObject("resource/stone", charGameObj.X(), charGameObj.Y(), 0.0, -1, nil)

		// check character has container
		putInContainer := false
		if (slots["back"] != nil) {
			// put log to container
			putInContainer = containers.Put(e, player, slots["back"].(string), stoneObj.Id(), -1)
		}

		// OR drop stone on the ground
		if !putInContainer {
			stoneObj.SetFloor(charGameObj.Floor())
			e.Floors()[stoneObj.Floor()].Insert(stoneObj)
			storage.GetClient().Updates <- stoneObj.Clone()
			e.SendResponseToVisionAreas(e.GameObjects()[player.CharacterGameObjectId], "add_object", map[string]interface{}{
				"object": stoneObj,
			})
		}

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
