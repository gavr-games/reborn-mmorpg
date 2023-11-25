package craft

import (
	"fmt"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects"
)

func Craft(e entity.IEngine, params map[string]interface{}) bool {
	playerId := int(params["playerId"].(float64))
	if player, ok := e.Players()[playerId]; ok {
		craftItemName := params["item_name"].(string)
		craftItemConfig := GetAtlas()[craftItemName].(map[string]interface{})
		charGameObj := e.GameObjects()[player.CharacterGameObjectId]
		slots := charGameObj.Properties()["slots"].(map[string]interface{})
	
		// Call check again to make sure nothing changed.
		// For example some player or mob could move to the place of future building
		if !Check(e, player, params) {
			return false
		}

		// Remove resources
		if (slots["back"] != nil) {
			container := e.GameObjects()[slots["back"].(string)]
			if !container.(entity.IContainerObject).RemoveItemsKinds(e, player, craftItemConfig["resources"].(map[string]interface{})) {
				e.SendSystemMessage("Cannot consume required resources.", player)
				return false
			}
		}

		// Create object
		if craftItemConfig["place_in_real_world"].(bool) { //create in the real world
			coords := params["inputs"].(map[string]interface{})["coordinates"].(map[string]interface{})
			x := coords["x"].(float64)
			y := coords["y"].(float64)
			rotation := params["inputs"].(map[string]interface{})["rotation"].(float64)
			itemObj, err := game_objects.CreateFromTemplate(e, craftItemName, x, y, 0.0)
			itemObj.Properties()["crafted_by_character_id"] = charGameObj.Id()
			itemObj.Rotate(rotation)
			if err != nil {
				e.SendSystemMessage(err.Error(), player)
				return false
			}
			e.GameObjects()[itemObj.Id()] = itemObj
			itemObj.SetFloor(charGameObj.Floor())
			e.Floors()[itemObj.Floor()].Insert(itemObj)
			storage.GetClient().Updates <- itemObj.Clone()

			e.SendResponseToVisionAreas(e.GameObjects()[player.CharacterGameObjectId], "add_object", map[string]interface{}{
				"object": itemObj,
			})
		} else {
			itemObj, err := game_objects.CreateFromTemplate(e, craftItemName, charGameObj.X(), charGameObj.Y(), 0.0)
			if err != nil {
				e.SendSystemMessage(err.Error(), player)
				return false
			}
			e.GameObjects()[itemObj.Id()] = itemObj

			// check character has container
			putInContainer := false
			if (slots["back"] != nil) {
				// put item to container
				containerTo := e.GameObjects()[slots["back"].(string)]
				putInContainer = containerTo.(entity.IContainerObject).Put(e, player, itemObj.Id(), -1)
			}

			// OR drop items on the ground
			if !putInContainer {
				itemObj.SetFloor(charGameObj.Floor())
				e.Floors()[itemObj.Floor()].Insert(itemObj)
				storage.GetClient().Updates <- itemObj.Clone()
				e.SendResponseToVisionAreas(e.GameObjects()[player.CharacterGameObjectId], "add_object", map[string]interface{}{
					"object": itemObj,
				})
			}
		}

		e.SendSystemMessage(fmt.Sprintf("You crafted %s.", craftItemName), player)
		return true
	} else {
		return false
	}
}
