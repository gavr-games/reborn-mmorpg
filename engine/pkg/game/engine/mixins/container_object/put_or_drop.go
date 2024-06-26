package container_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (cont *ContainerObject) PutOrDrop(e entity.IEngine, charGameObj entity.IGameObject, itemId string, position int) bool {
	playerId := charGameObj.GetProperty("player_id").(int)
	if player, ok := e.Players().Load(playerId); ok {
		container := cont.gameObj
		if item, itemOk := e.GameObjects().Load(itemId); itemOk {
			if !container.(entity.IContainerObject).Put(e, player, itemId, position) {
				item.SetGameAreaId(charGameObj.GameAreaId())
				item.SetProperty("visible", true)
				if gameArea, gaOk := e.GameAreas().Load(item.GameAreaId()); gaOk {
					gameArea.Insert(item)
				}
				e.SendGameObjectUpdate(item, "add_object")
			}
		} else {
			return false
		}

		return true
	}

	return false
}
