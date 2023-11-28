package character_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/serializers"
)

func (obj *CharacterObject) SelectTarget(e entity.IEngine, targetId string) bool {
	target := e.GameObjects()[targetId]

	if playerId, found := obj.Properties()["player_id"]; found {
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

	if targetable, ok := target.Properties()["targetable"]; ok {
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
	if oldTargetId, ok := obj.Properties()["target_id"]; ok {
		if oldTargetId != nil {
			obj.(entity.ICharacterObject).DeselectTarget(e)
		}
	}

	obj.Properties()["target_id"] = target.Id()
	storage.GetClient().Updates <- obj.Clone()

	if player != nil {
		e.SendResponse("select_target", serializers.GetInfo(e.GameObjects(), target), player)
	}

	return true
}