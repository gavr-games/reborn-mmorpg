package mob_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (mob *MobObject) Follow(targetObj entity.IGameObject) {
	// Check only owner can ask mob to follow
	if playerId := targetObj.GetProperty("player_id"); playerId != nil {
		playerIdInt := playerId.(int)
		if player, ok := mob.Engine.Players().Load(playerIdInt); ok {
			if mob.GetProperty("owner_id") != nil && targetObj.Id() == mob.GetProperty("owner_id").(string) {
				mob.State = StartFollowState
				mob.TargetObjectId = targetObj.Id()
			} else {
				mob.Engine.SendSystemMessage("You are not the owner of this creature.", player)
			}
		}
	} else { // allow to follow other object for future (not only players)
		mob.State = StartFollowState
		mob.TargetObjectId = targetObj.Id()
	}
}
