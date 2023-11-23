package hatcheries

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/containers"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/claims"
)

func CheckHatch(e entity.IEngine, player *entity.Player, hatcheryId string) bool {
	hatchery := e.GameObjects()[hatcheryId]
	resources := hatchery.Properties()["hatching_resources"].(map[string]interface{})
	charGameObj := e.GameObjects()[player.CharacterGameObjectId]
	slots := charGameObj.Properties()["slots"].(map[string]interface{})

	if hatchery == nil {
		e.SendSystemMessage("Hatchery does not exist.", player)
		return false
	}

	// check object type
	if hatchery.Properties()["type"].(string) != "hatchery" {
		e.SendSystemMessage("Please choose hatchery.", player)
		return false
	}

	// Check claim access
	if !claims.CheckAccess(e, charGameObj, hatchery) {
		e.SendSystemMessage("You don't have an access to this claim.", player)
		return false
	}

	if hatchery.CurrentAction != nil {
		e.SendSystemMessage("Hatchery is already hatching.", player)
		return false
	}

	// Check near the hatchery
	if !hatchery.IsCloseTo(charGameObj) {
		e.SendSystemMessage("You need to be closer to the hatchery.", player)
		return false
	}

	// Check has resources
	if len(resources) != 0 {
		// check character has container
		if (slots["back"] == nil) {
			e.SendSystemMessage("You don't have container with required resources.", player)
			return false
		}
		// check container has items
		if !containers.HasItemsKinds(e, slots["back"].(string), resources) {
			e.SendSystemMessage("You don't have required resources.", player)
			return false
		}
	}

	// Remove resources
	if (slots["back"] != nil) {
		if !containers.RemoveItemsKinds(e, player, slots["back"].(string), resources) {
			e.SendSystemMessage("Cannot consume required resources.", player)
			return false
		}
	}

	e.SendSystemMessage("Hatching has started.", player)
	return true
}