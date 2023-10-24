package targets

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

func Deselect(e entity.IEngine, obj *entity.GameObject) bool {
	obj.Properties["target_id"] = nil
	storage.GetClient().Updates <- obj
	if playerId, found := obj.Properties["player_id"]; found {
		playerIdInt := playerId.(int)
		if player, ok := e.Players()[playerIdInt]; ok {
			e.SendResponse("deselect_target", map[string]interface{}{}, player)
		}
	}

	return true
}
