package containers

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects"
)

func GetItems(e entity.IEngine, containerId string) map[string]interface{} {
	container := e.GameObjects()[containerId]
	itemIds := container.Properties["items_ids"].([]interface{})
	cont := game_objects.GetInfo(e.GameObjects(), container)
	cont["items"] = make([]map[string]interface{}, len(itemIds))

  for i, itemId := range itemIds {
		if itemId != "" {
    	cont["items"].([]map[string]interface{})[i] = game_objects.GetInfo(e.GameObjects(), e.GameObjects()[itemId.(string)])
		}
  }
	return cont
}
