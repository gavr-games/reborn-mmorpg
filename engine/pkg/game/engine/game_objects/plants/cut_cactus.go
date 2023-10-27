package plants

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/containers"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects"
)

// This func is called via delayed action mechanism
// params: playerId, cactusId
func CutCactus(e entity.IEngine, params map[string]interface{}) bool {
	playerId := int(params["playerId"].(float64))
	if player, ok := e.Players()[playerId]; ok {
		cactusId := params["cactusId"].(string)
		cactus := e.GameObjects()[cactusId]
		charGameObj := e.GameObjects()[player.CharacterGameObjectId]
		slots := charGameObj.Properties["slots"].(map[string]interface{})

		if cactus == nil {
			e.SendSystemMessage("Cactus does not exist.", player)
			return false
		}

		// Create cactus slice
		sliceObj, err := game_objects.CreateFromTemplate("resource/cactus_slice", charGameObj.X, charGameObj.Y)
		if err != nil {
			e.SendSystemMessage(err.Error(), player)
			return false
		}
		e.GameObjects()[sliceObj.Id] = sliceObj

		// check character has container
		putInContainer := false
		if (slots["back"] != nil) {
			// put log to container
			putInContainer = containers.Put(e, player, slots["back"].(string), sliceObj.Id, -1)
		}

		// OR drop logs on the ground
		if !putInContainer {
			sliceObj.Floor = charGameObj.Floor
			e.Floors()[sliceObj.Floor].Insert(sliceObj)
			storage.GetClient().Updates <- sliceObj
			e.SendResponseToVisionAreas(e.GameObjects()[player.CharacterGameObjectId], "add_object", map[string]interface{}{
				"object": sliceObj,
			})
		}

		// Decrease slices stored in the cactus
		resources := cactus.Properties["resources"].(map[string]interface{})
		resources["cactus_slice"] = resources["cactus_slice"].(float64) - 1.0

		// Remove cactus if no cactus_slice inside
		if resources["cactus_slice"].(float64) <= 0 {
			e.SendGameObjectUpdate(cactus, "remove_object")

			e.Floors()[cactus.Floor].FilteredRemove(e.GameObjects()[cactusId], func(b utils.IBounds) bool {
				return cactusId == b.(*entity.GameObject).Id
			})
			e.GameObjects()[cactusId] = nil
			delete(e.GameObjects(), cactusId)
		} else {
			storage.GetClient().Updates <- cactus
		}

		e.SendSystemMessage("You received a cactus slice.", player)
	} else {
		return false
	}

	return true
}