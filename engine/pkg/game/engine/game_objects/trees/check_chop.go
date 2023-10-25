package trees

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/characters"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects"
)

func CheckChop(e entity.IEngine, player *entity.Player, treeId string) bool {
	tree := e.GameObjects()[treeId]
	charGameObj := e.GameObjects()[player.CharacterGameObjectId]

	if tree == nil {
		e.SendSystemMessage("Tree does not exist.", player)
		return false
	}

	// Check has Axe equipped
	if _, equipped := characters.HasTypeEquipped(e, charGameObj, "axe"); !equipped {
		e.SendSystemMessage("You need to equip axe.", player)
		return false
	}

	// Check near the tree
	if !game_objects.AreClose(tree, charGameObj) {
		e.SendSystemMessage("You need to be closer to the tree.", player)
		return false
	}

	return true
}