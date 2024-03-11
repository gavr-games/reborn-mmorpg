package hatchery_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/claims"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (hatchery *HatcheryObject) CheckHatch(e entity.IEngine, charGameObj entity.IGameObject) bool {
	resources := hatchery.Properties()["hatching_resources"].(map[string]interface{})
	slots := charGameObj.Properties()["slots"].(map[string]interface{})

	playerId := charGameObj.Properties()["player_id"].(int)
	player, ok := e.Players().Load(playerId)
	if player == nil || !ok {
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

	if hatchery.CurrentAction() != nil {
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
		if slots["back"] == nil {
			e.SendSystemMessage("You don't have container with required resources.", player)
			return false
		}

		var (
			container entity.IGameObject
			contOk bool
		)
		if container, contOk = e.GameObjects().Load(slots["back"].(string)); !contOk {
			return false
		}

		// check container has items
		if !container.(entity.IContainerObject).HasItemsKinds(e, resources) {
			e.SendSystemMessage("You don't have required resources.", player)
			return false
		}

		// Remove resources
		if !container.(entity.IContainerObject).RemoveItemsKinds(e, player, resources) {
			e.SendSystemMessage("Cannot consume required resources.", player)
			return false
		}
	}

	e.SendSystemMessage("Hatching has started.", player)
	return true
}
