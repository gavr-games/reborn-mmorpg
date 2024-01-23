package craft

import (
	"fmt"
	"slices"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/claims"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

const (
	CraftDistance = 0.5
)

func Check(e entity.IEngine, player *entity.Player, params map[string]interface{}, checkDistanceNow bool) bool {
	craftItemName := params["item_name"].(string)
	craftItemConfig := GetAtlas()[craftItemName].(map[string]interface{})
	charGameObj := e.GameObjects()[player.CharacterGameObjectId]
	slots := charGameObj.Properties()["slots"].(map[string]interface{})

	// Has required tools equipped
	requiredTools := craftItemConfig["tools"].([]string)
	for _, requiredTool := range requiredTools {
		if _, equipped := charGameObj.(entity.ICharacterObject).HasTypeEquipped(e, requiredTool); !equipped {
			e.SendSystemMessage(fmt.Sprintf("You need to equip %s.", requiredTool), player)
			return false
		}
	}

	// Check has resources
	if len(craftItemConfig["resources"].(map[string]interface{})) != 0 {
		// check character has container
		if slots["back"] == nil {
			e.SendSystemMessage("You don't have container with required resources.", player)
			return false
		}
		// check container has items
		container := e.GameObjects()[slots["back"].(string)]
		if !container.(entity.IContainerObject).HasItemsKinds(e, craftItemConfig["resources"].(map[string]interface{})) {
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
		rotation := params["inputs"].(map[string]interface{})["rotation"].(float64)
		tempGameObj, err := game_objects.CreateFromTemplate(e, craftItemName, x, y, 0.0)
		if err != nil {
			e.SendSystemMessage(err.Error(), player)
			return false
		}
		tempGameObj.SetFloor(charGameObj.Floor())
		tempGameObj.Rotate(rotation)

		// Check claim access
		if !claims.CheckAccess(e, charGameObj, tempGameObj) {
			e.SendSystemMessage("You don't have an access to this claim.", player)
			return false
		}

		// Check not intersecting with another objects
		possibleCollidableObjects := e.Floors()[charGameObj.Floor()].RetrieveIntersections(utils.Bounds{
			X:      tempGameObj.X(),
			Y:      tempGameObj.Y(),
			Width:  tempGameObj.Width(),
			Height: tempGameObj.Height(),
		})

		if len(possibleCollidableObjects) > 0 {
			for _, val := range possibleCollidableObjects {
				gameObj := val.(entity.IGameObject)
				if gameObj.Id() == charGameObj.Id() && tempGameObj.Properties()["collidable"].(bool) {
					e.SendSystemMessage("Cannot build it here. There is something in the way.", player)
					return false
				}
				if collidable, ok := gameObj.Properties()["collidable"]; ok {
					if collidable.(bool) {
						e.SendSystemMessage("Cannot build it here. There is something in the way.", player)
						return false
					}
				}
				// Check can build only on allowed surfaces
				if gameObj.Type() == "surface" {
					if !slices.Contains(craftItemConfig["surfaces"].([]string), gameObj.Kind()) {

						e.SendSystemMessage(fmt.Sprintf("Cannot build it on %s.", gameObj.Kind()), player)
						return false
					}
				}
			}
		}

		if tempGameObj.Kind() == "claim_obelisk" {
			// check cannot create 2 claims
			if claimId, hasClaimId := charGameObj.Properties()["claim_obelisk_id"]; hasClaimId {
				if claimId != nil {
					e.SendSystemMessage("Cannot build second claim.", player)
					return false
				}
			}
			// check cannot create if there is another claim area
			if len(possibleCollidableObjects) > 0 {
				for _, val := range possibleCollidableObjects {
					if val.(entity.IGameObject).Kind() == "claim_area" {
						e.SendSystemMessage("Cannot build it here. There is another claim area in the way.", player)
						return false
					}
				}
			}
		}

		if charGameObj.GetDistance(tempGameObj) > CraftDistance {
			if checkDistanceNow {
				e.SendSystemMessage("You need to be closer.", player)
				return false
			} else {
				// move to object to craft it
				charGameObj.SetMoveToCoordsByObject(tempGameObj)
			}
		}
	}

	return true
}
