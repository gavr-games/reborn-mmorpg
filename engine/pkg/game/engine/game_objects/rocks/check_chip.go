package rocks

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/characters"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects"
)

func CheckChip(e entity.IEngine, player *entity.Player, rockId string) bool {
	rock := e.GameObjects()[rockId]
	charGameObj := e.GameObjects()[player.CharacterGameObjectId]

	if rock == nil {
		e.SendSystemMessage("Rock does not exist.", player)
		return false
	}

	// Check has Pickaxe equipped
	if !characters.HasTypeEquipped(e, charGameObj, "pickaxe") {
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