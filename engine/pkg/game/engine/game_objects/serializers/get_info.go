package serializers

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// Finds information about game object and prepares it for serialization
func GetInfo(e entity.IEngine, obj entity.IGameObject) map[string]interface{} {
	info := obj.Properties()
	if obj.Kind() == "player" {
		// Inject slot items info
		for slotKey, itemId := range obj.GetProperty("slots").(map[string]interface{}) {
			if itemId != nil {
				if item, found := e.GameObjects().Load(itemId.(string)); found {
					info["slots"].(map[string]interface{})[slotKey] = GetInfo(e, item)
				}
			}
		}
		// Inject target info
		if targetId := obj.GetProperty("target_id"); targetId != nil {
			if target, found := e.GameObjects().Load(targetId.(string)); found {
				info["target"] = GetInfo(e, target)
			}
		}
	}

	// Inject crafted by info
	if craftedById := obj.GetProperty("crafted_by_character_id"); craftedById != nil {
		if owner, foundOwner := e.GameObjects().Load(craftedById.(string)); foundOwner {
			info["crafted_by"] = GetInfo(e, owner)
		}
	}
	return info
}
