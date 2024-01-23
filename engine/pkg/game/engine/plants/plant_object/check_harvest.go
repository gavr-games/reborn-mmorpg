package plant_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/claims"
)

func (plant *PlantObject) CheckHarvest(e entity.IEngine, charGameObj entity.IGameObject) bool {
	playerId := charGameObj.Properties()["player_id"].(int)
	player := e.Players()[playerId]
	if player == nil {
		return false
	}

	// check object type
	if plant.Properties()["type"].(string) != "plant" {
		e.SendSystemMessage("Please choose plant.", player)
		return false
	}

	// Check claim access
	if !claims.CheckAccess(e, charGameObj, plant) {
		e.SendSystemMessage("You don't have an access to this claim.", player)
		return false
	}

	// Check has Shovel equipped
	if _, equipped := charGameObj.(entity.ICharacterObject).HasTypeEquipped(e, "shovel"); !equipped {
		e.SendSystemMessage("You need to equip shovel.", player)
		return false
	}

	return true
}