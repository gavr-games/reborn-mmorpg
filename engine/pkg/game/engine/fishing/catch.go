package fishing

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// This func is called via delayed action mechanism
// params: characterId, rodId
func Catch(e entity.IEngine, params map[string]interface{}) bool {
	var (
		rod, character entity.IGameObject
		rodOk, charOk bool
	)
	if rod, rodOk = e.GameObjects().Load(params["rodId"].(string)); !rodOk {
		return false
	}
	if character, charOk = e.GameObjects().Load(params["characterId"].(string)); !charOk {
		return false
	}
	ok, _ := rod.(entity.IFishingRodObject).Catch(e, character)
	return ok
}