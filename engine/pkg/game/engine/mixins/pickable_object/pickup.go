package pickable_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

func (obj *PickableObject) Pickup(e entity.IEngine, player *entity.Player) bool {
	var (
		charGameObj, container entity.IGameObject
		charOk, contOk bool
	)
	item := obj.gameObj
	itemId := item.Id()
	if charGameObj, charOk = e.GameObjects().Load(player.CharacterGameObjectId); !charOk {
		return false
	}
	slots := charGameObj.GetProperty("slots").(map[string]interface{})

	// intersects with character
	itemBounds := utils.Bounds{
		X:      item.X(),
		Y:      item.Y(),
		Width:  item.Width(),
		Height: item.Height(),
	}
	if !charGameObj.Intersects(itemBounds) {
		e.SendSystemMessage("You are too far away.", player)
		return false
	}

	// not in another container
	containerId := item.GetProperty("container_id")
	if containerId != nil {
		e.SendSystemMessage("Item is already in another container.", player)
		return false
	}

	// check character has container
	if slots["back"] == nil {
		e.SendSystemMessage("You don't have container to put item to.", player)
		return false
	}

	// put to container
	if containerId == nil {
		if container, contOk = e.GameObjects().Load(slots["back"].(string)); !contOk {
			return false
		}
		if !container.(entity.IContainerObject).Put(e, player, itemId, -1) {
			return false
		}
	}

	// remove from world
	if _, itemStillExists := e.GameObjects().Load(itemId); itemStillExists { // if item was stackable it is destroyed in Put func
		if gameArea, gaOk := e.GameAreas().Load(item.GameAreaId()); gaOk {
			gameArea.FilteredRemove(item, func(b utils.IBounds) bool {
				return itemId == b.(entity.IGameObject).Id()
			})
		}
		item.SetProperty("visible", false)

		storage.GetClient().Updates <- item.Clone()
		e.SendResponseToVisionAreas(charGameObj, "remove_object", map[string]interface{}{
			"object": map[string]interface{}{
				"Id": itemId,
			},
		})
	}

	e.SendResponseToVisionAreas(charGameObj, "pickup_object", map[string]interface{}{
		"character_id": charGameObj.Id(),
		"id":           itemId,
	})

	return true
}
