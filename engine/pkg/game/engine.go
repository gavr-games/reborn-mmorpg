package game

import (
	"encoding/json"
	"fmt"

	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine"
)

const FloorSize = 1000.0
const FloorCount = 4
const InitialPlayerX = 4.0
const InitialPlayerY = 4.0

type EngineResponse struct {
	ResponseType string

	ResponseData map[string]interface{}
}

// Engine runs the game
type Engine struct {
	floors []*utils.Quadtree

	players map[int]*Player

	gameObjectsId int
	gameObjects map[int]*entity.GameObject

	// Inbound messages from the clients.
	commands chan *ClientCommand

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewEngine() *Engine {
	return &Engine{
		players:     make(map[int]*Player),
		gameObjectsId: 0, //TODO: migrate to UUID
		gameObjects: make(map[int]*entity.GameObject),
		floors:      make([]*utils.Quadtree, FloorCount),
		commands:    make(chan *ClientCommand),
		register:    make(chan *Client),
		unregister:  make(chan *Client),
	}
}

func (e *Engine) CreatePlayerVisionArea(player *Player) *entity.GameObject {
	charGameObj := e.gameObjects[player.CharacterGameObjectId]
	e.gameObjectsId++
	additionalProps := make(map[string]interface{})
	additionalProps["player_id"] = player.Id
	gameObj := engine.CreateGameObject("player_vision_area", e.gameObjectsId, charGameObj.X, charGameObj.Y, additionalProps)
	gameObj.Floor = 0
	e.gameObjects[e.gameObjectsId] = gameObj
	e.floors[0].Insert(gameObj)
	player.VisionAreaGameObjectId = gameObj.Id
	return gameObj
}

func (e *Engine) CreatePlayer(client *Client) *Player {
	e.gameObjectsId++
	player := &Player{
		Id: client.character.Id,
		CharacterGameObjectId: 0,
		VisionAreaGameObjectId: 0,
		Client: client,
		VisibleObjects: make([]int, 100),
	}
	e.players[player.Id] = player
	additionalProps := make(map[string]interface{})
	additionalProps["player_id"] = player.Id
	gameObj := engine.CreateGameObject("player", e.gameObjectsId, InitialPlayerX, InitialPlayerY, additionalProps)
	gameObj.Floor = 0
	e.gameObjects[e.gameObjectsId] = gameObj
	e.floors[0].Insert(gameObj)
	player.CharacterGameObjectId = gameObj.Id
	return player
}

func (e *Engine) RegisterClient(client *Client) {
	// lookup if engine has already created player object
	if player, ok := e.players[client.character.Id]; ok {
		if player.Client != nil {
			// close previous socket connection for this player
			close(player.Client.send)
		} else {
			e.CreatePlayerVisionArea(player)
			e.gameObjects[player.CharacterGameObjectId].Properties["visible"] = true
			player.VisibleObjects = make([]int, 100)
		}
		player.Client = client
	} else {
		player = e.CreatePlayer(client)
		e.CreatePlayerVisionArea(player)
	}
	if player, ok := e.players[client.character.Id]; ok {
		visionArea := e.gameObjects[player.VisionAreaGameObjectId]
		visibleObjects := e.floors[0].RetrieveIntersections(utils.Bounds{
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
					n++
				}
			} else {
				visibleObjects[n] = val
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
		case client.send <- message:
		default:
			//close(client.send)
			//delete(h.clients, client)
		}
	}
}

// Main engine loop
func (e *Engine) Run() {
	engine.LoadGameObjects(e.floors, e.gameObjects, &e.gameObjectsId, FloorSize)
	//log.Println("Objects count %v", e.gameObjects[1])
	for {
		select {
		case client := <-e.register:
			e.RegisterClient(client)
		case client := <-e.unregister:
			if player, ok := e.players[client.character.Id]; ok {
				close(client.send)
				player.Client = nil
				e.floors[0].FilteredRemove(e.gameObjects[player.VisionAreaGameObjectId], func(b utils.IBounds) bool {
					return player.VisionAreaGameObjectId == b.(entity.GameObject).Id
				})
				e.gameObjects[player.VisionAreaGameObjectId] = nil
				e.gameObjects[player.CharacterGameObjectId].Properties["visible"] = false
				player.VisibleObjects = nil
			}
		default:
			//main game loop
		/*case cmd := <-e.commands:
			for client := range e.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(e.clients, client)
				}
			}*/
		}
	}
}
