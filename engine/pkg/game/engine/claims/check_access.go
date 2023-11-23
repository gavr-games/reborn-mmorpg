package claims

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// Check that charGameObj can interract with targetObj
// it can interract if the targetObj is not inside the claim area of another player
func CheckAccess(e entity.IEngine, charGameObj entity.IGameObject, targetObj entity.IGameObject) bool {
	// Check not intersecting with claim areas
	possibleCollidableObjects := e.Floors()[targetObj.Floor()].RetrieveIntersections(utils.Bounds{
		X:      targetObj.X(),
		Y:      targetObj.Y(),
		Width:  targetObj.Width(),
		Height: targetObj.Height(),
	})

	if len(possibleCollidableObjects) > 0 {
		for _, val := range possibleCollidableObjects {
			obj := val.(entity.IGameObject)
			if obj.Properties()["kind"] == "claim_area" {
				obelisk := e.GameObjects()[obj.Properties()["claim_obelisk_id"].(string)]
				owner := e.GameObjects()[obelisk.Properties()["crafted_by_character_id"].(string)]
				return owner.Id() == charGameObj.Id()
			}
		}
	}

	return true
}