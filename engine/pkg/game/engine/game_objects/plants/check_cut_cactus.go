package plants

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/characters"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/claims"
)

func CheckCutCactus(e entity.IEngine, player *entity.Player, cactusId string) bool {
	cactus := e.GameObjects()[cactusId]
	charGameObj := e.GameObjects()[player.CharacterGameObjectId]

	if cactus == nil {
		e.SendSystemMessage("Cactus does not exist.", player)
		return false
	}

	// check object type
	if cactus.Properties["type"].(string) != "plant" {
		e.SendSystemMessage("Please choose plant.", player)
		return false
	}

	// Check claim access
	if !claims.CheckAccess(e, charGameObj, cactus) {
		e.SendSystemMessage("You don't have an access to this claim.", player)
		return false
	}

	// Check has Knife equipped
	if _, equipped := characters.HasTypeEquipped(e, charGameObj, "knife"); !equipped {
		e.SendSystemMessage("You need to equip knife.", player)
		return false
	}

	// Check near the cactus
	if !cactus.IsCloseTo(charGameObj) {
		e.SendSystemMessage("You need to be closer to the cactus.", player)
		return false
	}

	return true
}