package mob_object

import (
	"context"
	"fmt"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (mob *MobObject) OrderToFollow(targetObj entity.IGameObject) {
	ctx := context.WithValue(context.Background(), "targetObjId", targetObj.Id())
	// Check only owner can ask mob to follow
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

			mob.FSM.Event(ctx, "follow")
		}
	}
}
