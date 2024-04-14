package craft

import (
	"fmt"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

// This func is trigerred by delayed action mechanism
func Craft(e entity.IEngine, params map[string]interface{}) bool {
	var (
		charGameObj, container entity.IGameObject
		charOk, contOk bool
	)
	playerId := int(params["playerId"].(float64))
	if player, ok := e.Players().Load(playerId); ok {
		craftItemName := params["item_name"].(string)
		craftItemConfig := GetAtlas()[craftItemName].(map[string]interface{})
		if charGameObj, charOk = e.GameObjects().Load(player.CharacterGameObjectId); !charOk {
			return false
		}
		slots := charGameObj.GetProperty("slots").(map[string]interface{})

		// Call check again to make sure nothing changed.
		// For example some player or mob could move to the place of future building
		if !Check(e, player, params, true) {
			return false
		}

		if slots["back"] == nil {
			e.SendSystemMessage("You don't have container", player)
			return false
		}

		if container, contOk = e.GameObjects().Load(slots["back"].(string)); !contOk {
			return false
		}

		if !container.(entity.IContainerObject).RemoveItemsKinds(e, player, craftItemConfig["resources"].(map[string]interface{})) {
			e.SendSystemMessage("Cannot consume required resources.", player)
			return false
		}

		// Create object
		if craftItemConfig["place_in_real_world"].(bool) { //create in the real world
			coords := params["inputs"].(map[string]interface{})["coordinates"].(map[string]interface{})
			x := coords["x"].(float64)
			y := coords["y"].(float64)
			rotation := params["inputs"].(map[string]interface{})["rotation"].(float64)
			itemObj, err := game_objects.CreateFromTemplate(e, craftItemName, x, y, 0.0)
			if err != nil {
				e.SendSystemMessage(err.Error(), player)
				return false
			}
			itemObj.SetProperty("crafted_by_character_id", charGameObj.Id())
			itemObj.Rotate(rotation)
			itemObj.SetGameAreaId(charGameObj.GameAreaId())
			e.GameObjects().Store(itemObj.Id(), itemObj)
			if gameArea, gaOk := e.GameAreas().Load(itemObj.GameAreaId()); gaOk {
				gameArea.Insert(itemObj)
			}
			storage.GetClient().Updates <- itemObj.Clone()

			e.SendResponseToVisionAreas(charGameObj, "add_object", map[string]interface{}{
				"object": itemObj.Clone(),
			})
		} else {
			itemObj, err := game_objects.CreateFromTemplate(e, craftItemName, charGameObj.X(), charGameObj.Y(), 0.0)
			if err != nil {
				e.SendSystemMessage(err.Error(), player)
				return false
			}
			e.GameObjects().Store(itemObj.Id(), itemObj)

			// put item to container or drop it to the ground
			container.(entity.IContainerObject).PutOrDrop(e, charGameObj, itemObj.Id(), -1)
		}

		charGameObj.(entity.ILevelingObject).AddExperience(e, fmt.Sprintf("craft/%s", craftItemName))

		e.SendSystemMessage(fmt.Sprintf("You crafted %s.", craftItemName), player)
		return true
	} else {
		return false
	}
}
