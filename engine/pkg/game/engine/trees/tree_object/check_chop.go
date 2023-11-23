package tree_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/characters"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/claims"
)

func (tree *TreeObject) CheckChop(e entity.IEngine, charGameObj entity.IGameObject) bool {
	playerId := charGameObj.Properties()["player_id"].(int)
	player := e.Players()[playerId]
	if player == nil {
		return false
	}

	// check object type
	if tree.Properties()["type"].(string) != "tree" {
		e.SendSystemMessage("Please choose a tree.", player)
		return false
	}

	// Check claim access
	if !claims.CheckAccess(e, charGameObj, tree) {
		e.SendSystemMessage("You don't have an access to this claim.", player)
		return false
	}

	// Check has Axe equipped
	if _, equipped := characters.HasTypeEquipped(e, charGameObj, "axe"); !equipped {
		e.SendSystemMessage("You need to equip axe.", player)
		return false
	}

	// Check near the tree
	if !tree.IsCloseTo(charGameObj) {
		e.SendSystemMessage("You need to be closer to the tree.", player)
		return false
	}

	return true
}
