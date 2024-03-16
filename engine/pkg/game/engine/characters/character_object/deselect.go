package character_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

func (obj *CharacterObject) DeselectTarget(e entity.IEngine) bool {
	targetId := obj.GetProperty("target_id")

	if targetId == nil {
		return true
	}

	if playerId := obj.GetProperty("player_id"); playerId != nil {
		playerIdInt := playerId.(int)
		if player, ok := e.Players().Load(playerIdInt); ok {
			e.SendResponse("deselect_target", map[string]interface{}{
				"id": targetId,
			}, player)
		}
	}

	obj.SetProperty("target_id", nil)
	storage.GetClient().Updates <- obj.Clone()

	return true
}
