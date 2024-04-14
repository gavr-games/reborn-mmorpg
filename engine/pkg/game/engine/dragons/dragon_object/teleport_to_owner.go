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
				if alive := dragon.GetProperty("alive"); alive == nil || !alive.(bool) {
					dragon.Engine.SendSystemMessage("The dragon is dead, ressurect it first.", player)
					return
				}
				dragonClone := dragon.Clone()
				dragon.Engine.SendResponseToVisionAreas(dragonClone, "remove_object", map[string]interface{}{
					"object": dragonClone,
				})
				dragon.StopEverything()
				if gameArea, gaOk := dragon.Engine.GameAreas().Load(dragon.GameAreaId()); gaOk {
					gameArea.FilteredRemove(dragon, func(b utils.IBounds) bool {
						return dragon.Id() == b.(entity.IGameObject).Id()
					})
				}
				dragon.SetX(charGameObj.X())
				dragon.SetY(charGameObj.Y())
				dragon.SetGameAreaId(charGameObj.GameAreaId())
				if gameArea, gaOk := dragon.Engine.GameAreas().Load(charGameObj.GameAreaId()); gaOk {
					gameArea.Insert(dragon)
				}
				dragon.Engine.SendGameObjectUpdate(dragon, "update_object")
			} else {
				dragon.Engine.SendSystemMessage("You are not the owner of this creature.", player)
			}
		}
	}
}
