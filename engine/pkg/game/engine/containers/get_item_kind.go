package containers

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// Checks if container has specified itemKind and return first
func GetItemKind(e entity.IEngine, containerId string, itemKind string) entity.IGameObject {
	container := e.GameObjects()[containerId]
	itemIds := container.Properties()["items_ids"].([]interface{})

	//TODO: search inside sub containers
  for _, itemId := range itemIds {
		if itemId != nil {
			kind := e.GameObjects()[itemId.(string)].Properties()["kind"].(string)
    	if kind == itemKind {
				return e.GameObjects()[itemId.(string)]
			}
		}
  }
	return nil
}
