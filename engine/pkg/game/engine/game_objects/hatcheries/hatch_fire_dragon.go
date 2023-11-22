package hatcheries

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mobs"
)

// This func is called via delayed action mechanism
// params: hatcheryId
func HatchFireDragon(e entity.IEngine, params map[string]interface{}) bool {
	hatcheryId := params["hatcheryId"].(string)
	hatchery := e.GameObjects()[hatcheryId]

	if hatchery != nil {
		// Create dragon
		// TODO: can we use here CreateGameObject?
		dragon, err := game_objects.CreateFromTemplate("mob/fire_dragon", hatchery.X, hatchery.Y, 0.0)
		if err != nil {
			return false
		}
		e.GameObjects()[dragon.Id] = dragon
		dragon.Floor = hatchery.Floor
		e.Floors()[dragon.Floor].Insert(dragon)
		e.Mobs()[dragon.Id] = mobs.NewMob(e, dragon.Id)

		storage.GetClient().Updates <- dragon.Clone()
		e.SendResponseToVisionAreas(dragon, "add_object", map[string]interface{}{
			"object": dragon,
		})

		// Remove hatchery
		e.SendGameObjectUpdate(hatchery, "remove_object")
		e.Floors()[hatchery.Floor].FilteredRemove(e.GameObjects()[hatcheryId], func(b utils.IBounds) bool {
			return hatcheryId == b.(*entity.GameObject).Id
		})
		e.GameObjects()[hatcheryId] = nil
		delete(e.GameObjects(), hatcheryId)
	} else {
		return false
	}

	return true
}