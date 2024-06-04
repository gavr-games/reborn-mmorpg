package burning_object

import (
	"errors"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/claims"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/effects"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (obj *BurningObject) Extinguish(e entity.IEngine, player *entity.Player) (bool, error) {
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

	// Check burning
	state := gameObj.GetProperty("state")
	if state.(string) != "burning" {
		e.SendSystemMessage("Not burning.", player)
		return false, errors.New("not burning")
	}

	// Extinguish
	for effectId, effect := range gameObj.Effects() {
		if effect.(map[string]interface{})["group"].(string) == "burning" {
			effects.Remove(e, effectId, gameObj)
		}
	}
	gameObj.SetProperty("state", "extinguished")
	e.SendGameObjectUpdate(gameObj, "update_object")

	return true, nil
}
