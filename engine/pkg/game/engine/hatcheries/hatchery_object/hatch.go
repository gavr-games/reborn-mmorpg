package hatchery_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

func (hatchery *HatcheryObject) Hatch(e entity.IEngine, mobPath string) bool {
	// Create dragon
	dragon := e.CreateGameObject(mobPath, hatchery.X(), hatchery.Y(), 0.0, hatchery.Floor(), nil)

	e.SendResponseToVisionAreas(dragon, "add_object", map[string]interface{}{
		"object": dragon,
	})

	// Remove hatchery
	e.SendGameObjectUpdate(hatchery, "remove_object")
	e.Floors()[hatchery.Floor()].FilteredRemove(hatchery, func(b utils.IBounds) bool {
		return hatchery.Id() == b.(entity.IGameObject).Id()
	})
	e.GameObjects()[hatchery.Id()] = nil
	delete(e.GameObjects(), hatchery.Id())
	storage.GetClient().Deletes <- hatchery.Id()

	return true
}
