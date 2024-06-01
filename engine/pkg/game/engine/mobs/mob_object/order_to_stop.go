package mob_object

import (
	"fmt"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (mob *MobObject) OrderToStop(targetObj entity.IGameObject) {
	// Check only owner can ask mob to stop
	if playerId := targetObj.GetProperty("player_id"); playerId != nil {
		playerIdInt := playerId.(int)
		if player, ok := mob.Engine.Players().Load(playerIdInt); ok {
			// Check owner
			ownerId := mob.GetProperty("owner_id")
			if ownerId == nil || targetObj.Id() != ownerId.(string) {
				mob.Engine.SendSystemMessage(fmt.Sprintf("You are not the owner of this %s.", mob.Kind()), player)
				return
			}

			// Check alive
			alive := mob.GetProperty("alive")
			if alive != nil && !alive.(bool) {
				mob.Engine.SendSystemMessage(fmt.Sprintf("The %s is dead", mob.Kind()), player)
				return
			}

			// Check not too far away
			if mob.GetDistance(targetObj) > ControlRange || mob.GameAreaId() != targetObj.GameAreaId() {
				mob.Engine.SendSystemMessage(fmt.Sprintf("The %s is too far away", mob.Kind()), player)
				return
			}

			mob.StopEverything()
		}
	}
}
