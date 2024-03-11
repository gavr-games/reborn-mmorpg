package container_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// Checks if container has specified itemKind and return first
func (cont *ContainerObject) GetItemKind(e entity.IEngine, itemKind string) entity.IGameObject {
	container := cont.gameObj
	itemIds := container.Properties()["items_ids"].([]interface{})

	for _, itemId := range itemIds {
		if itemId != nil {
			if item, ok := e.GameObjects().Load(itemId.(string)); ok {
				if item.Kind() == itemKind {
					return item
				}
				// Search inside sub containers
				if item.Type() == "container" {
					res := item.(entity.IContainerObject).GetItemKind(e, itemKind)
					if res != nil {
						return res
					}
				}
			}
		}
	}
	return nil
}
