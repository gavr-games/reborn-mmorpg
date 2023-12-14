package container_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (cont *ContainerObject) AddResource(e entity.IEngine, player *entity.Player, resourceObj entity.IGameObject, amount float64) bool {
	container := cont.gameObj

	resourceStackable := false
	if value, ok := resourceObj.Properties()["stackable"]; ok {
		resourceStackable = value.(bool)
	}

	if resourceStackable {
		resourceExist := false
		for _, itemId := range container.Properties()["items_ids"].([]interface{}) {
			if itemId != nil {
				item := e.GameObjects()[itemId.(string)]
				if item.Kind() == resourceObj.Kind() {
					resourceExist = true
					item.Properties()["amount"] = item.Properties()["amount"].(float64) + amount
					e.SendGameObjectUpdate(item, "update_object")
					break
				}
			}
		}

		if resourceExist {
			return true
		}

		resourceObj.Properties()["amount"] = amount
		return container.(entity.IContainerObject).Put(e, player, resourceObj.Id(), -1)
	}

	if container.Properties()["free_capacity"].(float64) > amount {
		for i := 0; i < int(amount); i++ {
			container.(entity.IContainerObject).Put(e, player, resourceObj.Id(), -1)
		}

		return true
	}

	return false
}
