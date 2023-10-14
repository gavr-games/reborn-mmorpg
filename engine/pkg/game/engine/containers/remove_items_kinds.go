package containers

import (
	"slices"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// Removes specified items inside the container (with specified counts)
// items example - {"log": 1.0, "stone": 2.0}
func RemoveItemsKinds(e entity.IEngine, player *entity.Player, containerId string, items map[string]interface{}) bool {
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
				if Remove(e, player, containerId, itemId.(string)) {
					itemsCounts[itemKind] = itemsCounts[itemKind] - 1.0
					if itemsCounts[itemKind] == 0.0 {
						itemsKinds = slices.DeleteFunc(itemsKinds, func(kind string) bool {
							return kind == itemKind
						})
					}
					if allCountsZero(itemsCounts) {
						return true
					}
				} else {
					return false
				}
			}
		}
  }
	return false
}
