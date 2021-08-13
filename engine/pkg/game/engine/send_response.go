package engine

import (
	"encoding/json"
	"fmt"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func SendResponse(e IEngine, responseType string, responseData map[string]interface{}, player *entity.Player) {
	resp := EngineResponse{
		ResponseType: responseType,
		ResponseData: responseData,
	}
	message, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
		return
	}
	select {
	case player.Client.GetSendChannel() <- message:
	default:
		UnregisterClient(e, player.Client)
	}
}
