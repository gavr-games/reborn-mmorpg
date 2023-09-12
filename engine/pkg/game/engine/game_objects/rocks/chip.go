package rocks

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/containers"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects"
)

// This func is called via delayed action mechanism
// params: playerId, rockId
func Chip(e entity.IEngine, params map[string]interface{}) bool {
	playerId := int(params["playerId"].(float64))
	if player, ok := e.Players()[playerId]; ok {
		rockId := params["rockId"].(string)
		rock := e.GameObjects()[rockId]
		charGameObj := e.GameObjects()[player.CharacterGameObjectId]
		slots := charGameObj.Properties["slots"].(map[string]interface{})

		if rock == nil {
			e.SendSystemMessage("Rock does not exist.", player)
			return false
		}

		// Create log
		stoneObj, err := game_objects.CreateFromTemplate("resource/stone", charGameObj.X, charGameObj.Y)
		if err != nil {
			e.SendSystemMessage(err.Error(), player)
			return false
		}
		e.GameObjects()[stoneObj.Id] = stoneObj

		// check character has container
		putInContainer := false
		if (slots["back"] != nil) {
			// put log to container
			putInContainer = containers.Put(e, player, slots["back"].(string), stoneObj.Id, -1)
		}

		// OR drop stone on the ground
		if !putInContainer {
			stoneObj.Floor = charGameObj.Floor
			e.Floors()[stoneObj.Floor].Insert(stoneObj)
			storage.GetClient().Updates <- stoneObj
		}

		// Decrease stones stored in the rock
		resources := rock.Properties["resources"].(map[string]interface{})
		resources["stone"] = resources["stone"].(float64) - 1.0

		// Remove rock if no stones inside
		if resources["stone"].(float64) <= 0 {
			e.SendGameObjectUpdate(rock, "remove_object")

			e.Floors()[0].FilteredRemove(e.GameObjects()[rockId], func(b utils.IBounds) bool {
				return rockId == b.(*entity.GameObject).Id
			})
			e.GameObjects()[rockId] = nil

			storage.GetClient().Deletes <- rock
		} else {
			storage.GetClient().Updates <- rock
		}

		e.SendSystemMessage("You received a stone.", player)
	} else {
		return false
	}

	return true
}