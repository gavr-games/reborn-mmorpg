package liftable_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// This func is trigerred by delayed action mechanism
// params: itemId, characterId
func Lift(e entity.IEngine, params map[string]interface{}) bool {
	var (
		item, character entity.IGameObject
		itemOk, charOk bool
	)
	if item, itemOk = e.GameObjects().Load(params["itemId"].(string)); !itemOk {
		return false
	}
	if character, charOk = e.GameObjects().Load(params["characterId"].(string)); !charOk {
		return false
	}
	return item.(entity.ILiftableObject).Lift(e, character)
}

// This func is trigerred by delayed action mechanism
// params: itemId, x, y, rotation, characterId
func PutLifted(e entity.IEngine, params map[string]interface{}) bool {
	var (
		item, character entity.IGameObject
		itemOk, charOk bool
	)
	if item, itemOk = e.GameObjects().Load(params["itemId"].(string)); !itemOk {
		return false
	}
	if character, charOk = e.GameObjects().Load(params["characterId"].(string)); !charOk {
		return false
	}
	x := params["x"].(float64)
	y := params["y"].(float64)
	rotation := params["rotation"].(float64)
	return item.(entity.ILiftableObject).PutLifted(e, character, x, y, rotation)
}
