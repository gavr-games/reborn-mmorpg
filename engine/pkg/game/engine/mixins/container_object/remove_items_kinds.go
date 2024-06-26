package container_object

import (
	"slices"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
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
	var (
		item entity.IGameObject
		itemOk bool
	)
	itemIds := container.GetProperty("items_ids").([]interface{})

	for _, itemId := range itemIds {
		if itemId != nil {
			if item, itemOk = e.GameObjects().Load(itemId.(string)); !itemOk {
				return false
			}
			itemKind := item.Kind()
			itemStackable := false
			if stackable := item.GetProperty("stackable"); stackable != nil {
				itemStackable = stackable.(bool)
			}
			// If item stackable substract "amount", otherwise remove items as 1 per each game_object
			if slices.Contains(itemsKinds, itemKind) {
				performRemove := true
				if itemStackable {
					item.SetProperty("amount", item.GetProperty("amount").(float64) - itemsCounts[itemKind])
					e.SendGameObjectUpdate(item, "update_object")
					if item.GetProperty("amount").(float64) != 0.0 {
						performRemove = false
						itemsCounts[itemKind] = 0.0
					} else {
						itemsCounts[itemKind] = 1.0 // will be decreased during object removing
					}
				}
				if performRemove {
					if container.(entity.IContainerObject).Remove(e, player, itemId.(string)) {
						e.RemoveGameObject(item)
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
			if item, itemOk := e.GameObjects().Load(itemId.(string)); itemOk {
				if item.Type() == "container" {
					if removeItemsKinds(e, player, item, itemsCounts, itemsKinds) {
						return true
					}
				}
			}
		}
	}

	return false
}
