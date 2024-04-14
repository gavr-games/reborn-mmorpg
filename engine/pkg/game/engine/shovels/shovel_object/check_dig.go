package shovel_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

const (
	SmallSize = 0.00000001
)

func (shovel *ShovelObject) CheckDig(e entity.IEngine, charGameObj entity.IGameObject) bool {
	playerId := charGameObj.GetProperty("player_id").(int)
	if player, ok := e.Players().Load(playerId); ok {
		// Check shovel equipped
		if _, equipped := charGameObj.(entity.ICharacterObject).HasTypeEquipped(e, "shovel"); !equipped {
			e.SendSystemMessage("You need to equip shovel.", player)
			return false
		}

		// Check char is on the grass
		gameArea, gaOk := e.GameAreas().Load(charGameObj.GameAreaId())
		if !gaOk {
			return false
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
					return true
				}
			}
		}

		e.SendSystemMessage("You should stand on the grass.", player)
	} else {
		return false
	}

	return false
}
