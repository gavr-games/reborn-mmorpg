package burning_object

import (
	"errors"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/claims"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	uuid "github.com/satori/go.uuid"
)

func (obj *BurningObject) Burn(e entity.IEngine, player *entity.Player) (bool, error) {
	var (
		charGameObj entity.IGameObject
		charOk bool
	)
	gameObj := obj.gameObj
	if charGameObj, charOk = e.GameObjects().Load(player.CharacterGameObjectId); !charOk {
		return false, errors.New("character not found")
	}

	// Check character is close enough
	if !gameObj.IsCloseTo(charGameObj) {
		e.SendSystemMessage("You need to be closer.", player)
		return false, errors.New("character needs to be closer")
	}

	// Check claim access
	if !claims.CheckAccess(e, charGameObj, gameObj) {
		e.SendSystemMessage("You don't have an access to this claim.", player)
		return false, errors.New("no access")
	}

	// Check fuel
	fuel := gameObj.GetProperty("fuel")
	if fuel.(float64) <= 0.0 {
		e.SendSystemMessage("No fuel.", player)
		return false, errors.New("no fuel")
	}

	// Check burning
	state := gameObj.GetProperty("state")
	if state.(string) == "burning" {
		e.SendSystemMessage("Already burning.", player)
		return false, errors.New("already burning")
	}

	// Burn
	effectId := uuid.NewV4().String()
	burningEffectMap := map[string]interface{}{
		"type":             "periodic",
		"attribute":        "fuel",
		"value":            -5000.0,
		"cooldown":         5000.0,
		"current_cooldown": 0.0,
		"number":           -1.0,
		"remove_on_zero":   true,
		"cant_go_negative": true,
		"finish_state":     "extinguished",
		"group":            "burning",
	}
	gameObj.SetEffect(effectId, burningEffectMap)
	effectMap := utils.CopyMap(burningEffectMap)
	effectMap["id"] = effectId
	effectMap["target_id"] = gameObj.Id()
	e.Effects().Store(effectId, effectMap)
	gameObj.SetProperty("state", "burning")
	e.SendGameObjectUpdate(gameObj, "update_object")

	return true, nil
}
