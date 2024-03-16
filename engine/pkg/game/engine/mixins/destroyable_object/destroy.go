package destroyable_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

func (obj *DestroyableObject) Destroy(e entity.IEngine, player *entity.Player) bool {
	var (
		charGameObj, container entity.IGameObject
		charOk, contOk bool
	)
	item := obj.gameObj
	if charGameObj, charOk = e.GameObjects().Load(player.CharacterGameObjectId); !charOk {
		return false
	}
	slots := charGameObj.GetProperty("slots").(map[string]interface{})

	// check equipped
	for _, slotItemId := range slots {
		if slotItemId == item.Id() {
			e.SendSystemMessage("Cannot destroy equipped item.", player)
			return false
		}
	}

	// check container belongs to character
	
	if containerId := item.GetProperty("container_id"); containerId != nil {
		if container, contOk = e.GameObjects().Load(containerId.(string)); !contOk {
			return false
		}
		if !container.(entity.IContainerObject).CheckAccess(e, player) {
			e.SendSystemMessage("You don't have access to this container", player)
			return false
		}
		if !container.(entity.IContainerObject).Remove(e, player, item.Id()) {
			return false
		}
	}

	// Destroy items inside container
	if item.Type() == "container" {
		itemIds := item.GetProperty("items_ids").([]interface{})
		for _, itemId := range itemIds {
			if itemId != nil {
				if itemInside, itemInsideOk := e.GameObjects().Load(itemId.(string)); itemInsideOk {
					itemInside.(entity.IDestroyableObject).Destroy(e, player)
				}
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
	e.GameObjects().Delete(item.Id())
	e.SendGameObjectUpdate(item, "remove_object")

	return true
}
