package destroyable_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (obj *DestroyableObject) Destroy(e entity.IEngine, player *entity.Player) bool {
	item := obj.gameObj
	charGameObj := e.GameObjects()[player.CharacterGameObjectId]
	slots := charGameObj.Properties()["slots"].(map[string]interface{})

	// check equipped
	for _, slotItemId := range slots {
		if slotItemId == item.Id() {
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
		if !container.(entity.IContainerObject).Remove(e, player, item.Id()) {
			return false
		}
	}

	// Destroy items inside container
	if (item.Type() == "container") {
		itemIds := item.Properties()["items_ids"].([]interface{})
		for _, itemId := range itemIds {
			if itemId != nil {
				e.GameObjects()[itemId.(string)].(entity.IDestroyableObject).Destroy(e, player)
			}
		}
	}

	// Destroy item
	// TODO: refactor to RemoveObject func in engine
	if item.Floor() != -1 {
		e.Floors()[item.Floor()].FilteredRemove(item, func(b utils.IBounds) bool {
			return item.Id() == b.(entity.IGameObject).Id()
		})
	}
	e.GameObjects()[item.Id()] = nil
	delete(e.GameObjects(), item.Id())
	e.SendGameObjectUpdate(item, "remove_object")

	return true
}
