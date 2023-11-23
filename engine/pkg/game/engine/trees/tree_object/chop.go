package tree_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/containers"
)

func (tree *TreeObject) Chop(e entity.IEngine, charGameObj entity.IGameObject) bool {
	playerId := charGameObj.Properties()["player_id"].(int)
	if player, ok := e.Players()[playerId]; ok {
		slots := charGameObj.Properties()["slots"].(map[string]interface{})

		// Create log
		logObj := e.CreateGameObject("resource/log", charGameObj.X(), charGameObj.Y(), 0.0, -1, nil)

		// check character has container
		putInContainer := false
		if (slots["back"] != nil) {
			// put log to container
			putInContainer = containers.Put(e, player, slots["back"].(string), logObj.Id(), -1)
		}

		// OR drop logs on the ground
		if !putInContainer {
			logObj.SetFloor(charGameObj.Floor())
			e.Floors()[logObj.Floor()].Insert(logObj)
			storage.GetClient().Updates <- logObj.Clone()
			e.SendResponseToVisionAreas(e.GameObjects()[player.CharacterGameObjectId], "add_object", map[string]interface{}{
				"object": logObj,
			})
		}

		// Decrease logs stored in the tree
		resources := tree.Properties()["resources"].(map[string]interface{})
		resources["log"] = resources["log"].(float64) - 1.0

		// Remove tree if no logs inside
		if resources["log"].(float64) <= 0 {
			e.SendGameObjectUpdate(tree, "remove_object")

			e.Floors()[tree.Floor()].FilteredRemove(e.GameObjects()[tree.Id()], func(b utils.IBounds) bool {
				return tree.Id() == b.(entity.IGameObject).Id()
			})
			e.GameObjects()[tree.Id()] = nil
			delete(e.GameObjects(), tree.Id())
		} else {
			storage.GetClient().Updates <- tree.Clone()
		}

		e.SendSystemMessage("You received a log.", player)
	} else {
		return false
	}

	return true
}
