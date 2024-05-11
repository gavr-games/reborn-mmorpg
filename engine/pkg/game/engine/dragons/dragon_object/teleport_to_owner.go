package dragon_object

import (
	"errors"

	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (dragon *DragonObject) TeleportToOwner(charGameObj entity.IGameObject) (bool, error) {
	// Check only owner can ask dragon to teleport
	if playerId := charGameObj.GetProperty("player_id"); playerId != nil {
		playerIdInt := playerId.(int)
		if player, ok := dragon.Engine.Players().Load(playerIdInt); ok {
			ownerId := dragon.GetProperty("owner_id")
			if ownerId != nil && charGameObj.Id() == ownerId.(string) {
				if charGameObj.GetProperty("current_dungeon_id") != nil {
					dragon.Engine.SendSystemMessage("You can't teleport in dungeon.", player)
					return false, errors.New("Can't teleport in dungeon")
				}
				if alive := dragon.GetProperty("alive"); alive == nil || !alive.(bool) {
					dragon.Engine.SendSystemMessage("The dragon is dead, ressurect it first.", player)
					return false, errors.New("The dragon is dead")
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

				return true, nil
			} else {
				dragon.Engine.SendSystemMessage("You are not the owner of this creature.", player)
				return false, errors.New("The dragon does not belong to player")
			}
		}
	}
	return false, errors.New("Player does not exist")
}
