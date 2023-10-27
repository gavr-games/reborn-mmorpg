package plants

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/characters"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects"
)

func CheckCutCactus(e entity.IEngine, player *entity.Player, cactusId string) bool {
	cactus := e.GameObjects()[cactusId]
	charGameObj := e.GameObjects()[player.CharacterGameObjectId]

	if cactus == nil {
		e.SendSystemMessage("Cactus does not exist.", player)
		return false
	}

	// Check has Knife equipped
	if _, equipped := characters.HasTypeEquipped(e, charGameObj, "knife"); !equipped {
		e.SendSystemMessage("You need to equip knife.", player)
		return false
	}

	// Check near the cactus
	if !game_objects.AreClose(cactus, charGameObj) {
		e.SendSystemMessage("You need to be closer to the cactus.", player)
		return false
	}

	return true
}