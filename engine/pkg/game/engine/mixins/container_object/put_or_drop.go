package container_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (cont *ContainerObject) PutOrDrop(e entity.IEngine, player *entity.Player, itemId string, position int) bool {
	container := cont.gameObj
	item := e.GameObjects()[itemId]

	if !container.(entity.IContainerObject).Put(e, player, itemId, position) {
		item.SetFloor(0)
		item.Properties()["visible"] = true
		e.Floors()[item.Floor()].Insert(item)
		e.SendGameObjectUpdate(item, "add_object")
		e.SendResponseToVisionAreas(e.GameObjects()[player.CharacterGameObjectId], "add_object", map[string]interface{}{
			"object": item,
		})
	}

	return true
}
