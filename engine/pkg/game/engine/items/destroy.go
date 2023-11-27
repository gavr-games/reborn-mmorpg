package items

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

func Destroy(e entity.IEngine, itemId string, player *entity.Player) bool {
	item := e.GameObjects()[itemId]
	charGameObj := e.GameObjects()[player.CharacterGameObjectId]
	slots := charGameObj.Properties()["slots"].(map[string]interface{})

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
	if (item.Properties()["container_id"] != nil) {
		container := e.GameObjects()[item.Properties()["container_id"].(string)]
		if !container.(entity.IContainerObject).CheckAccess(e, player) {
			e.SendSystemMessage("You don't have access to this container", player)
			return false
		}
		if !container.(entity.IContainerObject).Remove(e, player, itemId) {
			return false
		}
	}

	// Destroy item
	e.GameObjects()[itemId] = nil
	delete(e.GameObjects(), itemId)
	storage.GetClient().Deletes <- itemId

	return true
}
