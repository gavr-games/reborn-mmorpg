package cactus_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/claims"
)

func (cactus *CactusObject) CheckCut(e entity.IEngine, charGameObj entity.IGameObject) bool {
	playerId := charGameObj.Properties()["player_id"].(int)
	player := e.Players()[playerId]
	if player == nil {
		return false
	}

	// check object type
	if cactus.Properties()["type"].(string) != "plant" {
		e.SendSystemMessage("Please choose plant.", player)
		return false
	}

	// Check claim access
	if !claims.CheckAccess(e, charGameObj, cactus) {
		e.SendSystemMessage("You don't have an access to this claim.", player)
		return false
	}

	// Check has Knife equipped
	if _, equipped := charGameObj.(entity.ICharacterObject).HasTypeEquipped(e, "knife"); !equipped {
		e.SendSystemMessage("You need to equip knife.", player)
		return false
	}

	return true
}