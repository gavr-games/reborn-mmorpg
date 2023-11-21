package rocks

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/characters"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/claims"
)

func CheckChip(e entity.IEngine, player *entity.Player, rockId string) bool {
	rock := e.GameObjects()[rockId]
	charGameObj := e.GameObjects()[player.CharacterGameObjectId]

	if rock == nil {
		e.SendSystemMessage("Rock does not exist.", player)
		return false
	}

	// check object type
	if rock.Properties["type"].(string) != "rock" {
		e.SendSystemMessage("Please choose rock.", player)
		return false
	}

	// Check claim access
	if !claims.CheckAccess(e, charGameObj, rock) {
		e.SendSystemMessage("You don't have an access to this claim.", player)
		return false
	}

	// Check has Pickaxe equipped
	if _, equipped := characters.HasTypeEquipped(e, charGameObj, "pickaxe"); !equipped {
		e.SendSystemMessage("You need to equip pickaxe.", player)
		return false
	}

	// Check near the rock
	if !game_objects.AreClose(rock, charGameObj) {
		e.SendSystemMessage("You need to be closer to the rock.", player)
		return false
	}

	return true
}