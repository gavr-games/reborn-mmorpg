package engine

import (
	"encoding/json"
	"fmt"

	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

const InitialPlayerX = 4.0
const InitialPlayerY = 4.0

func CreatePlayer(e IEngine, client entity.IClient) *entity.Player {
	player := &entity.Player{
		Id: client.GetCharacter().Id,
		CharacterGameObjectId: "",
		VisionAreaGameObjectId: "",
		Client: client,
		VisibleObjects: make([]string, 100, 10000),
	}
	e.Players()[player.Id] = player
	additionalProps := make(map[string]interface{})
	additionalProps["player_id"] = player.Id
	gameObj := CreateGameObject("player", InitialPlayerX, InitialPlayerY, additionalProps)
	gameObj.Floor = 0
	e.GameObjects()[gameObj.Id] = gameObj
	e.Floors()[gameObj.Floor].Insert(gameObj)
	player.CharacterGameObjectId = gameObj.Id
	return player
}

func CreatePlayerVisionArea(e IEngine, player *entity.Player) *entity.GameObject {
	charGameObj := e.GameObjects()[player.CharacterGameObjectId]
	additionalProps := make(map[string]interface{})
	additionalProps["player_id"] = player.Id
	gameObj := CreateGameObject("player_vision_area", charGameObj.X, charGameObj.Y, additionalProps)
	gameObj.Floor = 0
	e.GameObjects()[gameObj.Id] = gameObj
	e.Floors()[gameObj.Floor].Insert(gameObj)
	player.VisionAreaGameObjectId = gameObj.Id
	return gameObj
}

// Process when new player logs into the game
func RegisterClient(e IEngine, client entity.IClient) {
	// lookup if engine has already created player object
	if player, ok := e.Players()[client.GetCharacter().Id]; ok {
		if player.Client != nil {
			// close previous socket connection for this player
			close(player.Client.GetSendChannel())
		} else {
			CreatePlayerVisionArea(e, player)
			e.GameObjects()[player.CharacterGameObjectId].Properties["visible"] = true
			player.VisibleObjects = make([]string, 100, 10000)
		}
		player.Client = client
	} else {
		player = CreatePlayer(e, client)
		CreatePlayerVisionArea(e, player)
	}
	if player, ok := e.Players()[client.GetCharacter().Id]; ok {
		visionArea := e.GameObjects()[player.VisionAreaGameObjectId]
		visibleObjects := e.Floors()[0].RetrieveIntersections(utils.Bounds{
			X:      visionArea.X,
			Y:      visionArea.Y,
			Width:  visionArea.Width,
			Height: visionArea.Height,
		})
		// Filter visible objects
		n := 0
    for _, val := range visibleObjects {
			if visible, ok := val.(*entity.GameObject).Properties["visible"]; ok {
				if visible.(bool) {
					visibleObjects[n] = val
					player.VisibleObjects = append(player.VisibleObjects, val.(*entity.GameObject).Id) //TODO: append performance
					n++
				}
			} else {
				visibleObjects[n] = val
				player.VisibleObjects = append(player.VisibleObjects, val.(*entity.GameObject).Id)
				n++
			}
    }
    visibleObjects = visibleObjects[:n]
		//Send json with VisibleObjects from vision area
		resp := EngineResponse{
			ResponseType: "init_game",
			ResponseData: map[string]interface{}{
				"visible_objects": visibleObjects,
			},
		}
		message, err := json.Marshal(resp)
    if err != nil {
        fmt.Println(err)
        return
    }
		select {
		case client.GetSendChannel() <- message:
		default:
			UnregisterClient(e, client)
		}
		SendGameObjectUpdate(e, e.GameObjects()[player.CharacterGameObjectId], "add_object")
	}
}