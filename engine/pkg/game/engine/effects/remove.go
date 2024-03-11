package effects

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// Removes effect from obj and engine map
func Remove(e entity.IEngine, effectId string, obj entity.IGameObject) {
	if obj != nil {
		obj.Effects()[effectId] = nil
		delete(obj.Effects(), effectId)
		e.SendGameObjectUpdate(obj, "update_object")
	}
	e.Effects().Delete(effectId)
}
