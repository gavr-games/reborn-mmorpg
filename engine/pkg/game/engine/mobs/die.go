package mobs

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func Die(e entity.IEngine, mobId string) bool {
	if _, ok := e.Mobs()[mobId]; ok {
		e.Mobs()[mobId] = nil
		delete(e.Mobs(), mobId)

		// remove from world
		mobObj := e.GameObjects()[mobId]
		e.Floors()[mobObj.Floor].FilteredRemove(e.GameObjects()[mobId], func(b utils.IBounds) bool {
			return mobId == b.(*entity.GameObject).Id
		})
		e.GameObjects()[mobId] = nil
		delete(e.GameObjects(), mobId)

		e.SendGameObjectUpdate(mobObj, "remove_object")
	}

	return true
}