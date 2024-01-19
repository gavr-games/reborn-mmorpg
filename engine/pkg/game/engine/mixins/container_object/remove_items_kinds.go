package container_object

import (
	"slices"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

// Removes specified items inside the container (with specified counts)
// items example - {"log": 1.0, "stone": 2.0}
func (cont *ContainerObject) RemoveItemsKinds(e entity.IEngine, player *entity.Player, items map[string]interface{}) bool {
	container := cont.gameObj

	itemsCounts := make(map[string]float64)
	var itemsKinds []string
	for k, v := range items {
		itemsCounts[k] = v.(float64)
		itemsKinds = append(itemsKinds, k)
	}

	return removeItemsKinds(e, player, container, itemsCounts, itemsKinds)
}

func removeItemsKinds(e entity.IEngine, player *entity.Player, container entity.IGameObject, itemsCounts map[string]float64, itemsKinds []string) bool {
	itemIds := container.Properties()["items_ids"].([]interface{})

	for _, itemId := range itemIds {
		if itemId != nil {
			item := e.GameObjects()[itemId.(string)]
			itemKind := item.Kind()
			itemStackable := false
			if value, ok := item.Properties()["stackable"]; ok {
				itemStackable = value.(bool)
			}
			// If item stackable substract "amount", otherwise remove items as 1 per each game_object
			if slices.Contains(itemsKinds, itemKind) {
				performRemove := true
				if itemStackable {
					item.Properties()["amount"] = item.Properties()["amount"].(float64) - itemsCounts[itemKind]
					e.SendGameObjectUpdate(item, "update_object")
					if item.Properties()["amount"].(float64) != 0.0 {
						performRemove = false
						itemsCounts[itemKind] = 0.0
					} else {
						itemsCounts[itemKind] = 1.0 // will be decreased during object removing
					}
				}
				if performRemove {
					if container.(entity.IContainerObject).Remove(e, player, itemId.(string)) {
						e.GameObjects()[item.Id()] = nil
						delete(e.GameObjects(), item.Id())
						storage.GetClient().Deletes <- item.Id()
						itemsCounts[itemKind] = itemsCounts[itemKind] - 1.0
					} else {
						return false
					}
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
				if removeItemsKinds(e, player, item, itemsCounts, itemsKinds) {
					return true
				}
			}
		}
	}

	return false
}
