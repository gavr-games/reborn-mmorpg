package character_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/serializers"
)

func (obj *CharacterObject) SelectTarget(e entity.IEngine, targetId string) bool {
	if target, targetOk := e.GameObjects().Load(targetId); targetOk {
		if playerId := obj.GetProperty("player_id"); playerId != nil {
			playerIdInt := playerId.(int)
			if player, ok := e.Players().Load(playerIdInt); ok {
				return selectTarget(e, obj, target, player)
			} else {
				return false
			}
		} else {
			return selectTarget(e, obj, target, nil)
		}
	} else {
		return false
	}
}

func selectTarget(e entity.IEngine, obj entity.IGameObject, target entity.IGameObject, player *entity.Player) bool {
	if target == nil {
		if player != nil {
			e.SendSystemMessage("Target does not exist.", player)
		}
		return false
	}

	if target.Id() == obj.Id() {
		if player != nil {
			e.SendSystemMessage("Cannot target self.", player)
		}
		return false
	}

	if targetable := target.GetProperty("targetable"); targetable != nil {
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

	// deselect previous target
	if oldTargetId := obj.GetProperty("target_id"); oldTargetId != nil {
		obj.(entity.ICharacterObject).DeselectTarget(e)
	}

	obj.SetProperty("target_id", target.Id())
	storage.GetClient().Updates <- obj.Clone()

	if player != nil {
		e.SendResponse("select_target", serializers.GetInfo(e, target), player)
	}

	return true
}