package pickable_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

func (obj *PickableObject) Pickup(e entity.IEngine, player *entity.Player) bool {
	item := obj.gameObj
	charGameObj := e.GameObjects()[player.CharacterGameObjectId]
	slots := charGameObj.Properties()["slots"].(map[string]interface{})

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
	if item.Properties()["container_id"] != nil {
		e.SendSystemMessage("Item is already in another container.", player)
		return false
	}

	// check character has container
	if slots["back"] == nil {
		e.SendSystemMessage("You don't have container to put item to.", player)
		return false
	}

	// put to container
	if item.Properties()["container_id"] == nil {
		container := e.GameObjects()[slots["back"].(string)]
		if !container.(entity.IContainerObject).Put(e, player, item.Id(), -1) {
			return false
		}
	}

	// remove from world
	e.Floors()[item.Floor()].FilteredRemove(item, func(b utils.IBounds) bool {
		return item.Id() == b.(entity.IGameObject).Id()
	})
	item.Properties()["visible"] = false

	storage.GetClient().Updates <- item.Clone()

	e.SendResponseToVisionAreas(charGameObj, "pickup_object", map[string]interface{}{
		"character_id": charGameObj.Id(),
		"id":           item.Id(),
	})
	e.SendResponseToVisionAreas(charGameObj, "remove_object", map[string]interface{}{
		"object": map[string]interface{}{
			"Id": item.Id(),
		},
	})

	return true
}
