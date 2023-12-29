package container_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (cont *ContainerObject) PutOrDrop(e entity.IEngine, charGameObj entity.IGameObject, itemId string, position int) bool {
	playerId := charGameObj.Properties()["player_id"].(int)
	if player, ok := e.Players()[playerId]; ok {
		container := cont.gameObj
		item := e.GameObjects()[itemId]

		if !container.(entity.IContainerObject).Put(e, player, itemId, position) {
			item.SetFloor(charGameObj.Floor())
			item.Properties()["visible"] = true
			e.Floors()[item.Floor()].Insert(item)
			e.SendGameObjectUpdate(item, "add_object")
			e.SendResponseToVisionAreas(e.GameObjects()[player.CharacterGameObjectId], "add_object", map[string]interface{}{
				"object": item,
			})
		}

		return true
	}

	return false
}
