package engine

import (
	"encoding/json"
	"fmt"

	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// send updates to all players who can see it
func SendResponseToVisionAreas(e IEngine, gameObj *entity.GameObject, responseType string, responseData map[string]interface{}) {
	intersectingObjects := e.Floors()[gameObj.Floor].RetrieveIntersections(utils.Bounds{
		X:      gameObj.X,
		Y:      gameObj.Y,
		Width:  gameObj.Width,
		Height: gameObj.Height,
	})
	resp := EngineResponse{
		ResponseType: responseType,
		ResponseData: responseData,
	}
	message, err := json.Marshal(resp)
	if err != nil {
			fmt.Println(err)
			return
	}
	for _, obj := range intersectingObjects {
		if obj.(*entity.GameObject).Type == "player" && obj.(*entity.GameObject).Properties["kind"].(string) != "player_vision_area" {
			playerId := obj.(*entity.GameObject).Properties["player_id"].(int)
			if player, ok := e.Players()[playerId]; ok {
				if player.Client != nil {
					select {
					case player.Client.GetSendChannel() <- message:
					default:
						UnregisterClient(e, player.Client)
					}
				}
			}
		}
	}
}
