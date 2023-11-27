package building_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/claims"
)

func (building *BuildingObject) Destroy(e entity.IEngine, player *entity.Player) bool {
	buildingObj := building.gameObj
	charGameObj := e.GameObjects()[player.CharacterGameObjectId]

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
	e.SendGameObjectUpdate(buildingObj, "remove_object")
	e.Floors()[buildingObj.Floor()].FilteredRemove(e.GameObjects()[buildingObj.Id()], func(b utils.IBounds) bool {
		return buildingObj.Id() == b.(entity.IGameObject).Id()
	})

	// Destroy building
	e.GameObjects()[buildingObj.Id()] = nil
	delete(e.GameObjects(), buildingObj.Id())
	storage.GetClient().Deletes <- buildingObj.Id()

	return true
}
