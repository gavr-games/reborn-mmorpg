package dragon_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

func (dragon *DragonObject) Die() {
	dragon.MobObject.Die()

	if ownerId := dragon.GetProperty("owner_id"); ownerId != nil {
		dragon.SetProperty("alive", false)
		dragon.SetProperty("visible", false)
		dragon.SetProperty("last_death", float64(dragon.Engine.CurrentTickTime()))

		dragon.Engine.GameObjects().Store(dragon.Id(), dragon) // put back dead dragon
		storage.GetClient().Updates <- dragon.Clone()
	}
}
