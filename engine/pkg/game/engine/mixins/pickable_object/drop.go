package pickable_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

func (obj *PickableObject) Drop(e entity.IEngine, player *entity.Player) bool {
	var (
		charGameObj, container entity.IGameObject
		charOk, contOk bool
	)
	item := obj.gameObj
	if charGameObj, charOk = e.GameObjects().Load(player.CharacterGameObjectId); !charOk {
		return false
	}
	slots := charGameObj.GetProperty("slots").(map[string]interface{})

	// check equipped
	for _, slotItemId := range slots {
		if slotItemId == item.Id() {
			e.SendSystemMessage("Cannot drop equipped item.", player)
			return false
		}
	}
	
	// check container belongs to character
	if containerId := item.GetProperty("container_id"); containerId != nil {
		if container, contOk = e.GameObjects().Load(containerId.(string)); !contOk {
			return false
		}
		if !container.(entity.IContainerObject).CheckAccess(e, player) {
			e.SendSystemMessage("You don't have access to this container", player)
			return false
		}
		//Remove from container
		if !container.(entity.IContainerObject).Remove(e, player, item.Id()) {
			return false
		}
	} else {
		e.SendSystemMessage("You can drop items only from container", player)
		return false
	}

	// Drop into the world
	item.SetGameAreaId(charGameObj.GameAreaId())
	item.SetProperty("visible", true)
	item.SetX(charGameObj.X())
	item.SetY(charGameObj.Y())
	if gameArea, gaOk := e.GameAreas().Load(item.GameAreaId()); gaOk {
		gameArea.Insert(item)
	}

	storage.GetClient().Updates <- item.Clone()

	e.SendResponseToVisionAreas(charGameObj, "add_object", map[string]interface{}{
		"object": item.Clone(),
	})

	return true
}
