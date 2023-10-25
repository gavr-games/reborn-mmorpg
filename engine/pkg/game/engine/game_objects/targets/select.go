package targets

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/serializers"
)

func Select(e entity.IEngine, obj *entity.GameObject, targetId string) bool {
	target := e.GameObjects()[targetId]

	if playerId, found := obj.Properties["player_id"]; found {
		playerIdInt := playerId.(int)
		if player, ok := e.Players()[playerIdInt]; ok {
			return selectTarget(e, obj, target, player)
		} else {
			return false
		}
	} else {
		return selectTarget(e, obj, target, nil)
	}
}

func selectTarget(e entity.IEngine, obj *entity.GameObject, target *entity.GameObject, player *entity.Player) bool {
	if target == nil {
		if player != nil {
			e.SendSystemMessage("Target does not exist.", player)
		}
		return false
	}

	if target.Id == obj.Id {
		if player != nil {
			e.SendSystemMessage("Cannot target self.", player)
		}
		return false
	}

	if targetable, ok := target.Properties["targetable"]; ok {
		if !targetable.(bool) {
			if player != nil {
				e.SendSystemMessage("Cannot target this object.", player)
			}
			return false
		}
	} else {
		if player != nil {
			e.SendSystemMessage("Cannot target this object.", player)
		}
		return false
	}

	obj.Properties["target_id"] = target.Id
	storage.GetClient().Updates <- obj

	if player != nil {
		e.SendResponse("select_target", serializers.GetInfo(e.GameObjects(), target), player)
	}

	return true
}