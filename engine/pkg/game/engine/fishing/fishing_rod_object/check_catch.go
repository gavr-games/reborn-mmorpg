package fishing_rod_object

import (
	"errors"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

const (
	CatchDistance = 0.5
)

func (fishingRod *FishingRodObject) CheckCatch(e entity.IEngine, charGameObj entity.IGameObject) (bool, error) {
	playerId := charGameObj.GetProperty("player_id").(int)
	if player, ok := e.Players().Load(playerId); ok {
		// Check fishing rod is equipped
		if _, equipped := charGameObj.(entity.ICharacterObject).HasTypeEquipped(e, "fishing_rod"); !equipped {
			e.SendSystemMessage("You need to equip fishing rod.", player)
			return false, errors.New("fishing rod is not equipped")
		}

		// Check char is near the water
		gameArea, gaOk := e.GameAreas().Load(charGameObj.GameAreaId())
		if !gaOk {
			return false, errors.New("game area not found")
		}
		possibleCollidableObjects := gameArea.RetrieveIntersections(utils.Bounds{
			X:      charGameObj.X() - CatchDistance,
			Y:      charGameObj.Y() - CatchDistance,
			Width:  charGameObj.Width() + CatchDistance * 2.0,
			Height: charGameObj.Height() + CatchDistance * 2.0,
		})

		if len(possibleCollidableObjects) > 0 {
			for _, val := range possibleCollidableObjects {
				obj := val.(entity.IGameObject)
				if obj.Kind() == "water" {
					return true, nil
				}
			}
		}

		e.SendSystemMessage("You should stand near the water.", player)
		return false, errors.New("character is not near the water")
	} else {
		return false, errors.New("player not found")
	}
}
