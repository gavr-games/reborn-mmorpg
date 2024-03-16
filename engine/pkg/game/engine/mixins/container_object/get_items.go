package container_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/serializers"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (cont *ContainerObject) GetItems(e entity.IEngine) map[string]interface{} {
	container := cont.gameObj
	itemIds := container.GetProperty("items_ids").([]interface{})
	contInfo := serializers.GetInfo(e, container)
	contInfo["items"] = make([]map[string]interface{}, len(itemIds))

	for i, itemId := range itemIds {
		if itemId != nil {
			if item, itemOk := e.GameObjects().Load(itemId.(string)); itemOk {
				contInfo["items"].([]map[string]interface{})[i] = serializers.GetInfo(e, item)
			}
		}
	}
	return contInfo
}
