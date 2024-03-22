package dragon_object

import (
	"slices"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

func (dragon *DragonObject) Release(charGameObj entity.IGameObject) {
	// Check only owner can release dragon
	if playerId := charGameObj.GetProperty("player_id"); playerId != nil {
		playerIdInt := playerId.(int)
		if player, ok := dragon.Engine.Players().Load(playerIdInt); ok {
			if dragon.GetProperty("owner_id") != nil && charGameObj.Id() == dragon.GetProperty("owner_id").(string) {
				// Check dragon is dead and remove completely
				if alive := dragon.GetProperty("alive"); alive == nil || !alive.(bool) {
					dragon.Engine.GameObjects().Delete(dragon.Id())
					storage.GetClient().Deletes <- dragon.Id()
				} else { // just allow dragon to become a free mob
					dragon.SetProperty("owner_id", nil)
					dragon.Engine.SendGameObjectUpdate(dragon, "update_object")
				}
				dragonIds := charGameObj.GetProperty("dragons_ids").([]interface{})
				index := slices.IndexFunc(dragonIds, func(id interface{}) bool { return id != nil && id.(string) == dragon.Id() })
				charGameObj.SetProperty("dragons_ids", append(dragonIds[:index], dragonIds[index+1:]...)) // remove dragon id from slice
				storage.GetClient().Updates <- charGameObj.Clone()
				dragon.Engine.SendResponse("dragons_info", charGameObj.(entity.ICharacterObject).GetDragonsInfo(dragon.Engine), player)
			} else {
				dragon.Engine.SendSystemMessage("You are not the owner of this creature.", player)
			}
		}
	}
}
