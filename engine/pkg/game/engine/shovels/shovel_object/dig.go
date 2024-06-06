package shovel_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

func (shovel *ShovelObject) Dig(e entity.IEngine, charGameObj entity.IGameObject) bool {
	playerId := charGameObj.GetProperty("player_id").(int)
	if player, ok := e.Players().Load(playerId); ok {
		// Check again
		if !shovel.CheckDig(e, charGameObj) {
			return false
		}

		grass := findGrass(e, charGameObj)

		// Add dirt
		dirt := e.CreateGameObject("surface/dirt", grass.X(), grass.Y(), 0.0, grass.GameAreaId(), nil)
		e.SendGameObjectUpdate(dirt, "add_object")

		// Remove grass
		e.RemoveGameObject(grass)

		e.SendSystemMessage("You've created some dirt.", player)
	} else {
		return false
	}

	return true
}

func findGrass(e entity.IEngine, charGameObj entity.IGameObject) entity.IGameObject {
	gameArea, gaOk := e.GameAreas().Load(charGameObj.GameAreaId())
	if !gaOk {
		return nil
	}
	possibleCollidableObjects := gameArea.RetrieveIntersections(utils.Bounds{
		X:      charGameObj.X() + charGameObj.Width()/2,
		Y:      charGameObj.Y() + charGameObj.Height()/2,
		Width:  SmallSize, // small size is used to determine the exact surface char is standing on
		Height: SmallSize, // if we use full char width the char could stand on multiple surffaces at once
	})

	if len(possibleCollidableObjects) > 0 {
		for _, val := range possibleCollidableObjects {
			obj := val.(entity.IGameObject)
			if obj.Kind() == "grass" {
				return obj
			}
		}
	}

	return nil
}
