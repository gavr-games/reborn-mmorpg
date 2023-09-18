package containers

import (
	"slices"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func allCountsZero(itemsCounts map[string]float64) bool {
	allZero := true
	for _, itemCount := range itemsCounts {
		if itemCount > 0.0 {
			allZero = false
		}
	}
	return allZero
}

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
			itemKind := e.GameObjects()[itemId.(string)].Properties["kind"].(string)
    	if slices.Contains(itemsKinds, itemKind) {
				itemsCounts[itemKind] = itemsCounts[itemKind] - 1.0
				if allCountsZero(itemsCounts) {
					return true
				}
			}
		}
  }
	return false
}
