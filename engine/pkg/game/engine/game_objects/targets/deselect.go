package targets

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

func Deselect(e entity.IEngine, obj *entity.GameObject) bool {
	targetId, ok := obj.Properties["target_id"]
	if !ok {
		return true
	}

	if targetId == nil {
		return true
	}

	if playerId, found := obj.Properties["player_id"]; found {
		playerIdInt := playerId.(int)
		if player, ok := e.Players()[playerIdInt]; ok {
			e.SendResponse("deselect_target", map[string]interface{}{
				"id": targetId,
			}, player)
		}
	}

	obj.Properties["target_id"] = nil
	storage.GetClient().Updates <- obj.Clone()

	return true
}
