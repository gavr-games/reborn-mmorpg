package trees

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/containers"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects"
)

// This func is called via delayed action mechanism
// params: playerId, treeId
func Chop(e entity.IEngine, params map[string]interface{}) bool {
	playerId := int(params["playerId"].(float64))
	if player, ok := e.Players()[playerId]; ok {
		treeId := params["treeId"].(string)
		tree := e.GameObjects()[treeId]
		charGameObj := e.GameObjects()[player.CharacterGameObjectId]
		slots := charGameObj.Properties()["slots"].(map[string]interface{})

		if tree == nil {
			e.SendSystemMessage("Tree does not exist.", player)
			return false
		}

		// Create log
		logObj, err := game_objects.CreateFromTemplate("resource/log", charGameObj.X(), charGameObj.Y(), 0.0)
		if err != nil {
			e.SendSystemMessage(err.Error(), player)
			return false
		}
		e.GameObjects()[logObj.Id()] = logObj

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

			e.Floors()[tree.Floor()].FilteredRemove(e.GameObjects()[treeId], func(b utils.IBounds) bool {
				return treeId == b.(entity.IGameObject).Id()
			})
			e.GameObjects()[treeId] = nil
			delete(e.GameObjects(), treeId)
		} else {
			storage.GetClient().Updates <- tree.Clone()
		}

		e.SendSystemMessage("You received a log.", player)
	} else {
		return false
	}

	return true
}