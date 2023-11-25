package container_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/serializers"
)

func (cont *ContainerObject) GetItems(e entity.IEngine) map[string]interface{} {
	container := cont.gameObj
	itemIds := container.Properties()["items_ids"].([]interface{})
	contInfo := serializers.GetInfo(e.GameObjects(), container)
	contInfo["items"] = make([]map[string]interface{}, len(itemIds))

  for i, itemId := range itemIds {
		if itemId != nil {
    	contInfo["items"].([]map[string]interface{})[i] = serializers.GetInfo(e.GameObjects(), e.GameObjects()[itemId.(string)])
		}
  }
	return contInfo
}
