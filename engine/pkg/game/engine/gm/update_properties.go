package gm

import (
	"encoding/json"
	"fmt"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// Update properties
// params: object_id(string), update_props(json)
func UpdateProperties(e entity.IEngine, charGameObj entity.IGameObject, params map[string]interface{}) bool {
	playerId := charGameObj.GetProperty("player_id").(int)
	player, ok := e.Players().Load(playerId)
	if player == nil || !ok {
		return false
	}

	// check character is Game Master
	if !charGameObj.GetProperty("game_master").(bool) {
		e.SendSystemMessage("You are not a Game Master, cheater!", player)
		return false
	}

	// Update properties
	if params["object_id"] == nil {
		e.SendSystemMessage("Wrong object_id.", player)
		return false
	}
	var (
		gameObj entity.IGameObject
		goOk bool
	)
	objId := params["object_id"].(string)
	if gameObj, goOk = e.GameObjects().Load(objId); !goOk {
		e.SendSystemMessage("Wrong object_id.", player)
		return false
	}

	var updateProps map[string]interface{}
	err := json.Unmarshal([]byte(params["update_props"].(string)), &updateProps)
	if err != nil {
		e.SendSystemMessage("Wrong update properties format.", player)
		return false
	}

	for propKey, propVal := range updateProps {
		gameObj.SetProperty(propKey, propVal)
	}

	e.SendResponseToVisionAreas(gameObj, "update_object", map[string]interface{}{
		"object": gameObj.Clone(),
	})

	e.SendSystemMessage(fmt.Sprintf("You've updated %s properties.", gameObj.Kind()), player)

	return true
}