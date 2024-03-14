package claims

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
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
			if obj.Kind() == "claim_area" {
				var (
					obelisk, owner entity.IGameObject
					obeliskOk, ownerOk bool
				)
				if obelisk, obeliskOk = e.GameObjects().Load(obj.GetProperty("claim_obelisk_id").(string)); !obeliskOk {
					return false
				}
				if owner, ownerOk = e.GameObjects().Load(obelisk.GetProperty("crafted_by_character_id").(string)); !ownerOk {
					return false
				}
				return owner.Id() == charGameObj.Id()
			}
		}
	}

	return true
}
