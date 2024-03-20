package dragon_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (dragon *DragonObject) TeleportToOwner(charGameObj entity.IGameObject) {
	// Check only owner can ask dragon to teleport
	if playerId := charGameObj.GetProperty("player_id"); playerId != nil {
		playerIdInt := playerId.(int)
		if player, ok := dragon.Engine.Players().Load(playerIdInt); ok {
			if dragon.GetProperty("owner_id") != nil && charGameObj.Id() == dragon.GetProperty("owner_id").(string) {
				// TODO: Check dragon is dead
				dragon.StopEverything()
				dragonClone := dragon.Clone()
				dragon.Engine.SendResponseToVisionAreas(dragonClone, "remove_object", map[string]interface{}{
					"object": dragonClone,
				})
				dragon.Engine.Floors()[dragon.Floor()].FilteredRemove(dragon, func(b utils.IBounds) bool {
					return dragon.Id() == b.(entity.IGameObject).Id()
				})
				dragon.SetX(charGameObj.X())
				dragon.SetY(charGameObj.Y())
				dragon.SetFloor(charGameObj.Floor())
				dragon.Engine.Floors()[charGameObj.Floor()].Insert(dragon)
				dragon.Engine.SendGameObjectUpdate(dragon, "update_object")
			} else {
				dragon.Engine.SendSystemMessage("You are not the owner of this creature.", player)
			}
		}
	}
}
