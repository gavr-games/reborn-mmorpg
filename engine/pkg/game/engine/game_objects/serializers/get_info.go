package serializers

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

// Finds information about game object and prepares it for serialization
func GetInfo(gameObjects map[string]entity.IGameObject, obj entity.IGameObject) map[string]interface{} {
	info := utils.CopyMap(obj.Properties())
	if obj.Kind() == "player" {
		// Inject slot items info
		for slotKey, itemId := range obj.Properties()["slots"].(map[string]interface{}) {
			if itemId != nil {
				info["slots"].(map[string]interface{})[slotKey] = GetInfo(gameObjects, gameObjects[itemId.(string)])
			}
		}
		// Inject target info
		if targetId, ok := obj.Properties()["target_id"]; ok {
			if targetId != nil {
				if target, found := gameObjects[targetId.(string)]; found {
					info["target"] = GetInfo(gameObjects, target)
				}
			}
		}
	}

	// Inject crafted by info
	if craftedById, isCrafted := obj.Properties()["crafted_by_character_id"]; isCrafted {
		if craftedById != nil {
			if owner, foundOwner := gameObjects[craftedById.(string)]; foundOwner {
				info["crafted_by"] = GetInfo(gameObjects, owner)
			}
		}
	}
	return info
}
