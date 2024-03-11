package equipable_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/serializers"
)

func (obj *EquipableObject) Equip(e entity.IEngine, player *entity.Player) bool {
	var (
		charGameObj, container entity.IGameObject
		charOk, contOk bool
	)
	item := obj.gameObj
	if charGameObj, charOk = e.GameObjects().Load(player.CharacterGameObjectId); !charOk {
		return false
	}
	slots := charGameObj.Properties()["slots"].(map[string]interface{})
	targetSlots := item.Properties()["target_slots"].(map[string]interface{})

	if item == nil {
		e.SendSystemMessage("Wrong item.", player)
		return false
	}
	
	// check already equipped
	for _, slotItemId := range slots {
		if slotItemId == item.Id() {
			e.SendSystemMessage("This item is already equipped.", player)
			return false
		}
	}

	// check character has free slot
	freeTargetSlot := ""
	for targetSlotKey, _ := range targetSlots {
		if slots[targetSlotKey] == nil {
			freeTargetSlot = targetSlotKey
			break
		}
	}
	if freeTargetSlot == "" {
		e.SendSystemMessage("No free slots to equip item.", player)
		return false
	}

	//check in container
	if (item.Properties()["container_id"] == nil) {
		e.SendSystemMessage("First pickup item to equip it.", player)
		return false
	}

	// check container belongs to character
	if (item.Properties()["container_id"] != nil) {
		if container, contOk = e.GameObjects().Load(item.Properties()["container_id"].(string)); !contOk {
			return false
		}
		if !container.(entity.IContainerObject).CheckAccess(e, player) {
			e.SendSystemMessage("You don't have access to this container", player)
			return false
		}
		// remove from container if in container
		if !container.(entity.IContainerObject).Remove(e, player, item.Id()) {
			return false
		}
	}
	
	// Add to slot
	slots[freeTargetSlot] = item.Id()
	storage.GetClient().Updates <- charGameObj.Clone()
	
	e.SendResponseToVisionAreas(charGameObj, "equip_item", map[string]interface{}{
		"slot": freeTargetSlot,
		"character_id": player.CharacterGameObjectId,
		"item": serializers.GetInfo(e, item),
	})

	return true
}
