package craft

import (
	"fmt"

	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/characters"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/containers"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects"
)

func Check(e entity.IEngine, player *entity.Player, params map[string]interface{}) bool {
	craftItemName := params["item_name"].(string)
	craftItemConfig := GetAtlas()[craftItemName].(map[string]interface{})
	charGameObj := e.GameObjects()[player.CharacterGameObjectId]
	slots := charGameObj.Properties["slots"].(map[string]interface{})

	// Has required tools equipped
	requiredTools := craftItemConfig["tools"].([]string)
	for _, requiredTool := range requiredTools {
		if !characters.HasTypeEquipped(e, charGameObj, requiredTool) {
			e.SendSystemMessage(fmt.Sprintf("You need to equip %s.", requiredTool), player)
			return false
		}
	}

	// Check has resources
	if len(craftItemConfig["resources"].(map[string]interface{})) != 0 {
		// check character has container
		if (slots["back"] == nil) {
			e.SendSystemMessage("You don't have container with required resources.", player)
			return false
		}
		// check container has items
		if !containers.HasItemsKinds(e, slots["back"].(string), craftItemConfig["resources"].(map[string]interface{})) {
			e.SendSystemMessage("You don't have required resources.", player)
			return false
		}
	}

	// Check is near if object is crafted in real world
	if craftItemConfig["place_in_real_world"].(bool) {
		// create temporary game object
		coords := params["inputs"].(map[string]interface{})["coordinates"].(map[string]interface{})
		x := coords["x"].(float64)
		y := coords["y"].(float64)
		tempGameObj, err := game_objects.CreateFromTemplate(craftItemName, x, y)
		if err != nil {
			e.SendSystemMessage(err.Error(), player)
			return false
		}
		//TODO: take into account the rotation
		if !game_objects.AreClose(tempGameObj, charGameObj) {
			e.SendSystemMessage("You need to be closer.", player)
			return false
		}

		// Check not intersecting with another objects
		possibleCollidableObjects := e.Floors()[charGameObj.Floor].RetrieveIntersections(utils.Bounds{
			X:      tempGameObj.X,
			Y:      tempGameObj.Y,
			Width:  tempGameObj.Width,
			Height: tempGameObj.Height,
		})

		if len(possibleCollidableObjects) > 0 {
			e.SendSystemMessage("Cannot build it here. There is something in the way.", player)
			return false
		}
	}

	return true
}