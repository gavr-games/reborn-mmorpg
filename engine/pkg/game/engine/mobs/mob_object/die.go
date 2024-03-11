package mob_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (mob *MobObject) Die() {
	mob.drop()

	// remove from world
	mob.Engine.Floors()[mob.Floor()].FilteredRemove(mob, func(b utils.IBounds) bool {
			return mob.Id() == b.(entity.IGameObject).Id()
	})
	mob.Engine.GameObjects()[mob.Id()] = nil
	delete(mob.Engine.GameObjects(), mob.Id())

	mob.Engine.Mobs().Delete(mob.Id())

	mob.Engine.SendGameObjectUpdate(mob, "remove_object")
}