package tree_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

func (tree *TreeObject) Chop(e entity.IEngine, charGameObj entity.IGameObject) bool {
	playerId := charGameObj.GetProperty("player_id").(int)
	if player, ok := e.Players().Load(playerId); ok {
		slots := charGameObj.GetProperty("slots").(map[string]interface{})

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
		logObj := e.CreateGameObject("resource/log", charGameObj.X(), charGameObj.Y(), 0.0, "", nil)

		// Put to container or drop to the ground
		if container, contOk := e.GameObjects().Load(slots["back"].(string)); contOk {
			container.(entity.IContainerObject).PutOrDrop(e, charGameObj, logObj.Id(), -1)
		} else {
			return false
		}

		// Decrease logs stored in the tree
		resources := tree.GetProperty("resources").(map[string]interface{})
		resources["log"] = resources["log"].(float64) - 1.0
		tree.SetProperty("resources", resources)

		// Remove tree if no logs inside
		if resources["log"].(float64) <= 0 {
			e.RemoveGameObject(tree)
		} else {
			storage.GetClient().Updates <- tree.Clone()
		}

		charGameObj.(entity.ILevelingObject).AddExperience(e, "chop_tree")

		e.SendSystemMessage("You received a log.", player)
	} else {
		return false
	}

	return true
}
