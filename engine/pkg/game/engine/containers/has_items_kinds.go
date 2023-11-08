package containers

import (
	"slices"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// Checks if container has specified items inside of it (with specified counts)
// items example - {"log": 1.0, "stone": 2.0}
func HasItemsKinds(e entity.IEngine, containerId string, items map[string]interface{}) bool {
	container := e.GameObjects()[containerId]
	itemIds := container.Properties["items_ids"].([]interface{})

	itemsCounts := make(map[string]float64)
	var itemsKinds []string
	for k, v := range items {
		itemsCounts[k] = v.(float64)
		itemsKinds = append(itemsKinds, k)
	}

	//TODO: search inside sub containers
  for _, itemId := range itemIds {
		if itemId != nil {
			item := e.GameObjects()[itemId.(string)]
			itemKind := item.Properties["kind"].(string)
			itemStackable := false
			if value, ok := item.Properties["stackable"]; ok {
				itemStackable = value.(bool)
			}
			// If item stackable check item has enough "amount", otherwise count items as 1 per each game_object
    	if slices.Contains(itemsKinds, itemKind) {
				if itemStackable {
					if item.Properties["amount"].(float64) >= itemsCounts[itemKind] {
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
	return false
}
