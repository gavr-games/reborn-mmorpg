package gm

import (
	"encoding/json"
	"fmt"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// Create object
func CreateObject(e entity.IEngine, charGameObj entity.IGameObject, params map[string]interface{}) bool {
	playerId := charGameObj.Properties()["player_id"].(int)
	player := e.Players()[playerId]
	if player == nil {
		return false
	}

	// check character is Game Master
	if !charGameObj.Properties()["game_master"].(bool) {
		e.SendSystemMessage("You are not a Game Master, cheater!", player)
		return false
	}

	// Create object
	objPath := params["object_path"].(string)
	offsetX := params["offset_x"].(float64)
	offsetY := params["offset_y"].(float64)
	var additionalProps interface{}
	err := json.Unmarshal([]byte(params["additional_props"].(string)), &additionalProps)
	if err != nil {
		e.SendSystemMessage("Wrong additional properties format.", player)
		return false
	}

	gameObj := e.CreateGameObject(objPath, charGameObj.X() + offsetX, charGameObj.Y() + offsetY, 0.0, charGameObj.Floor(), additionalProps.(map[string]interface{}))

	e.SendResponseToVisionAreas(gameObj, "add_object", map[string]interface{}{
		"object": gameObj,
	})

	e.SendSystemMessage(fmt.Sprintf("You've created %s.", objPath), player)

	return true
}