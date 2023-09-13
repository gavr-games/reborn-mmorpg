package items

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/containers"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects"
)

func Destroy(e entity.IEngine, itemId string, player *entity.Player) bool {
	item := e.GameObjects()[itemId]
	charGameObj := e.GameObjects()[player.CharacterGameObjectId]
	slots := charGameObj.Properties["slots"].(map[string]interface{})

	if item == nil {
		e.SendSystemMessage("Wrong item.", player)
		return false
	}

	// check equipped
	for _, slotItemId := range slots {
		if slotItemId == itemId {
			e.SendSystemMessage("Cannot destroy equipped item.", player)
			return false
		}
	}

	// check container belongs to character
	if (item.Properties["container_id"] != nil) {
		if !containers.CheckAccess(e, player, e.GameObjects()[item.Properties["container_id"].(string)]) {
			e.SendSystemMessage("You don't have access to this container", player)
			return false
		}
		if !containers.Remove(e, player, item.Properties["container_id"].(string), itemId) {
			return false
		}
	} else { //destroy item in the world
		// Check near item
		if !game_objects.AreClose(item, charGameObj) {
			e.SendSystemMessage("You need to be closer to the item.", player)
			return false
		}
		e.SendGameObjectUpdate(item, "remove_object")
		e.Floors()[item.Floor].FilteredRemove(e.GameObjects()[itemId], func(b utils.IBounds) bool {
			return itemId == b.(*entity.GameObject).Id
		})
	}

	// Destroy item
	e.GameObjects()[itemId] = nil
	storage.GetClient().Deletes <- item

	return true
}