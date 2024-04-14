package building_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/claims"
)

func (building *BuildingObject) Destroy(e entity.IEngine, player *entity.Player) bool {
	buildingObj := building.gameObj
	var (
		charGameObj entity.IGameObject
		charOk bool
	)
	if charGameObj, charOk = e.GameObjects().Load(player.CharacterGameObjectId); !charOk {
		return false
	}

	// Check claim access
	if !claims.CheckAccess(e, charGameObj, buildingObj) {
		e.SendSystemMessage("You don't have an access to this claim.", player)
		return false
	}

	// Check near building
	if !buildingObj.IsCloseTo(charGameObj) {
		e.SendSystemMessage("You need to be closer to the item.", player)
		return false
	}

	// Destroy building
	e.RemoveGameObject(buildingObj)

	return true
}
