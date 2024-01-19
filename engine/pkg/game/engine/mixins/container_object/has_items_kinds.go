package container_object

import (
	"slices"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// Checks if container has specified items inside of it (with specified counts)
// items example - {"log": 1.0, "stone": 2.0}
func (cont *ContainerObject) HasItemsKinds(e entity.IEngine, items map[string]interface{}) bool {
	container := cont.gameObj

	itemsCounts := make(map[string]float64)
	var itemsKinds []string
	for k, v := range items {
		itemsCounts[k] = v.(float64)
		itemsKinds = append(itemsKinds, k)
	}

	return calcItemsKinds(e, container, itemsCounts, itemsKinds)
}

func calcItemsKinds(e entity.IEngine, container entity.IGameObject, itemsCounts map[string]float64, itemsKinds []string) bool {
	itemIds := container.Properties()["items_ids"].([]interface{})

	for _, itemId := range itemIds {
		if itemId != nil {
			item := e.GameObjects()[itemId.(string)]
			itemKind := item.Kind()
			itemStackable := false
			if value, ok := item.Properties()["stackable"]; ok {
				itemStackable = value.(bool)
			}
			// If item stackable check item has enough "amount", otherwise count items as 1 per each game_object
			if slices.Contains(itemsKinds, itemKind) {
				if itemStackable {
					if item.Properties()["amount"].(float64) >= itemsCounts[itemKind] {
						itemsCounts[itemKind] = 0.0
					}
				} else {
					itemsCounts[itemKind] = itemsCounts[itemKind] - 1.0
				}
				if itemsCounts[itemKind] == 0.0 {
					itemsKinds = slices.DeleteFunc(itemsKinds, func(kind string) bool {
						return kind == itemKind
					})
				}
				if len(itemsKinds) == 0 {
					return true
				}
			}
		}
	}

	//Search inside sub containers
	for _, itemId := range itemIds {
		if itemId != nil {
			item := e.GameObjects()[itemId.(string)]
			if item.Type() == "container" {
				if calcItemsKinds(e, item, itemsCounts, itemsKinds) {
					return true
				}
			}
		}
	}

	return false
}
