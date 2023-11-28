package equipable_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/serializers"
)

func (obj *EquipableObject) Unequip(e entity.IEngine, player *entity.Player) bool {
	item := obj.gameObj
	charGameObj := e.GameObjects()[player.CharacterGameObjectId]
	slots := charGameObj.Properties()["slots"].(map[string]interface{})

	if item == nil {
		e.SendSystemMessage("Wrong item.", player)
		return false
	}
	
	// check equipped
	itemSlotKey := ""
	for key, slotItemId := range slots {
		if slotItemId == item.Id() {
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
	if (item.Properties()["container_id"] == nil) {
		container := e.GameObjects()[slots["back"].(string)]
		if !container.(entity.IContainerObject).Put(e, player, item.Id(), -1) {
			return false
		}
	}
	
	// Remove from slot
	slots[itemSlotKey] = nil
	storage.GetClient().Updates <- charGameObj.Clone()
	
	e.SendResponseToVisionAreas(e.GameObjects()[player.CharacterGameObjectId], "unequip_item", map[string]interface{}{
		"slot": itemSlotKey,
		"character_id": player.CharacterGameObjectId,
		"item": serializers.GetInfo(e.GameObjects(), item),
	})

	return true
}
