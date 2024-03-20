package leveling_object

import (
	"errors"
	"fmt"
	"math"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

const (
	EXP_AMOUNT_DIVIDER = 0.05
	EXP_INCREASE_POWER = 2.0
	MAX_DRAGONS_MULTIPLIER = 10
)

func (obj *LevelingObject) AddExperience(e entity.IEngine, action string) (bool, error) {
	gameObj := obj.gameObj
	amount, ok := GetAtlas()[action]

	if !ok {
		return false, errors.New("Action not found")
	}

	// Send exp message if player found
	if playerId := gameObj.GetProperty("player_id"); playerId != nil {
		if player, pOk := e.Players().Load(playerId.(int)); pOk {
			e.SendSystemMessage(fmt.Sprintf("You received %d exp.", int(amount)), player)
		}
	}

	return obj.updateExpAndLevel(e, amount)
}

func (obj *LevelingObject) updateExpAndLevel(e entity.IEngine, amount float64) (bool, error) {
	gameObj := obj.gameObj

	currentExp := gameObj.GetProperty("experience").(float64)
	newExp := amount + currentExp
	currentLevel := gameObj.GetProperty("level").(float64)
	nextLevelExp := getExpForNextLevel(currentLevel)

	if newExp < nextLevelExp {
		gameObj.SetProperty("experience", newExp)
		storage.GetClient().Updates <- gameObj.Clone()
		e.SendResponseToVisionAreas(gameObj, "set_exp", map[string]interface{}{
			"exp": newExp,
			"object_id": gameObj.Id(),
		})
		return true, nil
	} else {
		gameObj.SetProperty("level", currentLevel + 1)
		gameObj.SetProperty("experience", 0.0)
		gameObj.SetProperty("max_dragons", math.Ceil(math.Sqrt((currentLevel + 1) * MAX_DRAGONS_MULTIPLIER)))
		storage.GetClient().Updates <- gameObj.Clone()
		e.SendResponseToVisionAreas(gameObj, "set_level", map[string]interface{}{
			"level": currentLevel + 1,
			"object_id": gameObj.Id(),
		})
		return obj.updateExpAndLevel(e, newExp - nextLevelExp)
	}
}

// https://blog.jakelee.co.uk/converting-levels-into-xp-vice-versa/
func getExpForNextLevel(level float64) float64 {
	totalExpForLevel := math.Pow(level / EXP_AMOUNT_DIVIDER, EXP_INCREASE_POWER)
	totalExpForNextLevel := math.Pow((level + 1) / EXP_AMOUNT_DIVIDER, EXP_INCREASE_POWER)
	return totalExpForNextLevel - totalExpForLevel
}
