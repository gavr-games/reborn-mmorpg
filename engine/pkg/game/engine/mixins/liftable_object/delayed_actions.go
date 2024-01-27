package liftable_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// This func is trigerred by delayed action mechanism
// params: itemId, characterId
func Lift(e entity.IEngine, params map[string]interface{}) bool {
	item := e.GameObjects()[params["itemId"].(string)].(entity.ILiftableObject)
	character := e.GameObjects()[params["characterId"].(string)]
	return item.Lift(e, character)
}

// This func is trigerred by delayed action mechanism
// params: itemId, x, y, rotation, characterId
func PutLifted(e entity.IEngine, params map[string]interface{}) bool {
	item := e.GameObjects()[params["itemId"].(string)].(entity.ILiftableObject)
	character := e.GameObjects()[params["characterId"].(string)]
	x := params["x"].(float64)
	y := params["y"].(float64)
	rotation := params["rotation"].(float64)
	return item.PutLifted(e, character, x, y, rotation)
}
