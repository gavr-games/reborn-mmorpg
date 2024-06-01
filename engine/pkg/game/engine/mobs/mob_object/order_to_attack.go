package mob_object

import (
	"fmt"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (mob *MobObject) OrderToAttack(ownerObj entity.IGameObject) {
	// Check only owner can ask mob to attack
	if playerId := ownerObj.GetProperty("player_id"); playerId != nil {
		playerIdInt := playerId.(int)
		if player, ok := mob.Engine.Players().Load(playerIdInt); ok {
			// Check owner
			ownerId := mob.GetProperty("owner_id")
			if ownerId == nil || ownerObj.Id() != ownerId.(string) {
				mob.Engine.SendSystemMessage(fmt.Sprintf("You are not the owner of this %s.", mob.Kind()), player)
				return
			}

			// Check alive
			alive := mob.GetProperty("alive")
			if alive != nil && !alive.(bool) {
				mob.Engine.SendSystemMessage(fmt.Sprintf("The %s is dead", mob.Kind()), player)
				return
			}

			// Check not too far away from owner
			if mob.GetDistance(ownerObj) > ControlRange || mob.GameAreaId() != ownerObj.GameAreaId() {
				mob.Engine.SendSystemMessage(fmt.Sprintf("The %s is too far away", mob.Kind()), player)
				return
			}

			// Check target is valid
			targetObjId := ownerObj.GetProperty("target_id")
			if targetObjId == nil || targetObjId.(string) == mob.Id() {
				mob.Engine.SendSystemMessage(fmt.Sprintf("Wrong target for %s", mob.Kind()), player)
				return
			}
			targetObj, toOk := mob.Engine.GameObjects().Load(targetObjId.(string))
			if !toOk {
				mob.Engine.SendSystemMessage(fmt.Sprintf("Wrong target for %s", mob.Kind()), player)
				return
			}

			// Check target is too far away
			if mob.GetDistance(targetObj) > ControlRange || mob.GameAreaId() != ownerObj.GameAreaId() {
				mob.Engine.SendSystemMessage(fmt.Sprintf("The %s is too far away from target", mob.Kind()), player)
				return
			}

			mob.Attack(targetObjId.(string))
		}
	}
}
