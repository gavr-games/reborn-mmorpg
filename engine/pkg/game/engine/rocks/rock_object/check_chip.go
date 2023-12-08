package rock_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/claims"
)

func (rock *RockObject) CheckChip(e entity.IEngine, charGameObj entity.IGameObject) bool {
	playerId := charGameObj.Properties()["player_id"].(int)
	player := e.Players()[playerId]
	if player == nil {
		return false
	}

	// check object type
	if rock.Properties()["type"].(string) != "rock" {
		e.SendSystemMessage("Please choose rock.", player)
		return false
	}

	// Check claim access
	if !claims.CheckAccess(e, charGameObj, rock) {
		e.SendSystemMessage("You don't have an access to this claim.", player)
		return false
	}

	// Check has Pickaxe equipped
	if _, equipped := charGameObj.(entity.ICharacterObject).HasTypeEquipped(e, "pickaxe"); !equipped {
		e.SendSystemMessage("You need to equip pickaxe.", player)
		return false
	}

	return true
}
