package engine

import (
	"encoding/json"
	"fmt"

	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// send new state of the game object to all players who can see it
func SendGameObjectUpdate(e IEngine, gameObj *entity.GameObject, updateType string) {
	intersectingObjects := e.Floors()[gameObj.Floor].RetrieveIntersections(utils.Bounds{
		X:      gameObj.X,
		Y:      gameObj.Y,
		Width:  gameObj.Width,
		Height: gameObj.Height,
	})
	resp := EngineResponse{
		ResponseType: updateType,
		ResponseData: map[string]interface{}{
			"object": gameObj,
		},
	}
	message, err := json.Marshal(resp)
	if err != nil {
			fmt.Println(err)
			return
	}
	for _, obj := range intersectingObjects {
		if obj.(*entity.GameObject).Type == "player_vision_area" {
			playerId := obj.(*entity.GameObject).Properties["player_id"]
			if player, ok := e.Players()[playerId.(int)]; ok {
				select {
				case player.Client.GetSendChannel() <- message:
				default:
					UnregisterClient(e, player.Client)
				}
			}
		}
	}
}
