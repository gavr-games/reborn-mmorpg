package container_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// Checks if container has specified itemKind and return first
func (cont *ContainerObject) GetItemKind(e entity.IEngine, itemKind string) entity.IGameObject {
	container := cont.gameObj
	itemIds := container.Properties()["items_ids"].([]interface{})

	//TODO: search inside sub containers
	for _, itemId := range itemIds {
		if itemId != nil {
			kind := e.GameObjects()[itemId.(string)].Kind()
			if kind == itemKind {
				return e.GameObjects()[itemId.(string)]
			}
		}
	}
	return nil
}
