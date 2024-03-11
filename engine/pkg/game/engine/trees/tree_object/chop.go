package tree_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

func (tree *TreeObject) Chop(e entity.IEngine, charGameObj entity.IGameObject) bool {
	playerId := charGameObj.Properties()["player_id"].(int)
	if player, ok := e.Players().Load(playerId); ok {
		slots := charGameObj.Properties()["slots"].(map[string]interface{})

		if slots["back"] == nil {
			e.SendSystemMessage("You don't have container", player)
			return false
		}

		if !tree.CheckChop(e, charGameObj) {
			return false
		}

		// Check near the tree
		if !tree.IsCloseTo(charGameObj) {
			e.SendSystemMessage("You need to be closer to the tree.", player)
			return false
		}

		// Create log
		logObj := e.CreateGameObject("resource/log", charGameObj.X(), charGameObj.Y(), 0.0, -1, nil)

		// Put to container or drop to the ground
		container := e.GameObjects()[slots["back"].(string)]
		container.(entity.IContainerObject).PutOrDrop(e, charGameObj, logObj.Id(), -1)

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
