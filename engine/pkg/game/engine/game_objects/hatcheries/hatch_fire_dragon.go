package hatcheries

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects"
)

// This func is called via delayed action mechanism
// params: hatcheryId
func HatchFireDragon(e entity.IEngine, params map[string]interface{}) bool {
	hatcheryId := params["hatcheryId"].(string)
	hatchery := e.GameObjects()[hatcheryId]

	if hatchery != nil {
		// Create dragon
		dragon, err := game_objects.CreateFromTemplate("mob/fire_dragon", hatchery.X, hatchery.Y)
		if err != nil {
			return false
		}
		e.GameObjects()[dragon.Id] = dragon
		dragon.Floor = hatchery.Floor
		e.Floors()[dragon.Floor].Insert(dragon)

		storage.GetClient().Updates <- dragon
		e.SendResponseToVisionAreas(dragon, "add_object", map[string]interface{}{
			"object": dragon,
		})

		// Remove hatchery
		e.SendGameObjectUpdate(hatchery, "remove_object")
		e.Floors()[0].FilteredRemove(e.GameObjects()[hatcheryId], func(b utils.IBounds) bool {
			return hatcheryId == b.(*entity.GameObject).Id
		})
		e.GameObjects()[hatcheryId] = nil

		storage.GetClient().Deletes <- hatchery
	} else {
		return false
	}

	return true
}