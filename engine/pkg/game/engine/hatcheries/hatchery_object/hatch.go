package hatchery_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/claims"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

func (hatchery *HatcheryObject) Hatch(e entity.IEngine, mobPath string) bool {
	// Create dragon
	dragon := e.CreateGameObject(mobPath, hatchery.X(), hatchery.Y(), 0.0, hatchery.GameAreaId(), nil)

	// Check hatchery is on claim and player has free dragon capacity
	if obelisk := claims.GetClaimObelisk(e, hatchery); obelisk != nil {
		if owner, ownerOk := e.GameObjects().Load(obelisk.GetProperty("crafted_by_character_id").(string)); ownerOk {
			dragonIds := owner.GetProperty("dragons_ids").([]interface{})
			if len(dragonIds) < int(owner.GetProperty("max_dragons").(float64)) {
				dragon.SetProperty("owner_id", owner.Id())
				owner.SetProperty("dragons_ids", append(dragonIds, dragon.Id()))
				storage.GetClient().Updates <- owner.Clone()
			}
		}
	}

	e.SendGameObjectUpdate(dragon, "add_object")

	// Remove hatchery
	e.RemoveGameObject(hatchery)

	return true
}
