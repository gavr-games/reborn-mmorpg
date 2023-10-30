package items

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/containers"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/serializers"
)

func Unequip(e entity.IEngine, itemId string, player *entity.Player) bool {
	item := e.GameObjects()[itemId]
	charGameObj := e.GameObjects()[player.CharacterGameObjectId]
	slots := charGameObj.Properties["slots"].(map[string]interface{})

	if item == nil {
		e.SendSystemMessage("Wrong item.", player)
		return false
	}
	
	// check equipped
	itemSlotKey := ""
	for key, slotItemId := range slots {
		if slotItemId == itemId {
			itemSlotKey = key
		}
	}
	if itemSlotKey == "" {
		e.SendSystemMessage("Item is not equipped.", player)
		return false
	}

	// check character has container
	if (slots["back"] == nil) {
		e.SendSystemMessage("You don't have container to put item to.", player)
		return false
	}

	// put to container
	if (item.Properties["container_id"] == nil) {
		if !containers.Put(e, player, slots["back"].(string), itemId, -1) {
			return false
		}
	}
	
	// Remove from slot
	slots[itemSlotKey] = nil
	storage.GetClient().Updates <- game_objects.Clone(charGameObj)
	
	e.SendResponseToVisionAreas(e.GameObjects()[player.CharacterGameObjectId], "unequip_item", map[string]interface{}{
		"slot": itemSlotKey,
		"character_id": player.CharacterGameObjectId,
		"item": serializers.GetInfo(e.GameObjects(), item),
	})

	return true
}
