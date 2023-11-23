package containers

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/serializers"
)

func GetItems(e entity.IEngine, containerId string) map[string]interface{} {
	container := e.GameObjects()[containerId]
	itemIds := container.Properties()["items_ids"].([]interface{})
	cont := serializers.GetInfo(e.GameObjects(), container)
	cont["items"] = make([]map[string]interface{}, len(itemIds))

  for i, itemId := range itemIds {
		if itemId != nil {
    	cont["items"].([]map[string]interface{})[i] = serializers.GetInfo(e.GameObjects(), e.GameObjects()[itemId.(string)])
		}
  }
	return cont
}
