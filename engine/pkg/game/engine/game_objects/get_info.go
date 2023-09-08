package game_objects

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

// Finds information about game object and prepares it for serialization
func GetInfo(gameObjects map[string]*entity.GameObject, obj *entity.GameObject) map[string]interface{} {
	info := utils.CopyMap(obj.Properties)
	if (obj.Properties["kind"].(string) == "player") { // Add slot items info
		for slotKey, itemId := range obj.Properties["slots"].(map[string]interface{}) {
			if itemId != nil {
				info["slots"].(map[string]interface{})[slotKey] = GetInfo(gameObjects, gameObjects[itemId.(string)])
			}
		}
	}
	return info
}