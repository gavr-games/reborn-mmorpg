package claims

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

// Get claim for the specified gameObj if any
func GetClaimObelisk(e entity.IEngine, gameObj entity.IGameObject) entity.IGameObject {
	// Check not intersecting with claim areas
	gameArea, gaOk := e.GameAreas().Load(gameObj.GameAreaId())
	if !gaOk {
		return nil
	}
	possibleCollidableObjects := gameArea.RetrieveIntersections(utils.Bounds{
		X:      gameObj.X(),
		Y:      gameObj.Y(),
		Width:  gameObj.Width(),
		Height: gameObj.Height(),
	})

	if len(possibleCollidableObjects) > 0 {
		for _, val := range possibleCollidableObjects {
			obj := val.(entity.IGameObject)
			if obj.Kind() == "claim_area" {
				if obelisk, obeliskOk := e.GameObjects().Load(obj.GetProperty("claim_obelisk_id").(string)); obeliskOk {
					return obelisk
				} else {
					return nil
				}
			}
		}
	}

	return nil
}
