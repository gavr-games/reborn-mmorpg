package game

import (
	"encoding/json"
	"fmt"

	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

const (
	FloorSize = 1000.0
	FloorCount = 4
	TickSize = 10 // Game tick size in ms
)

// Engine runs the game
type Engine struct {
	tickTime int64 //last tick time in milliseconds
	floors []*utils.Quadtree // slice of global game areas, underground, etc
	players map[int]*entity.Player // map of all players
	gameObjects map[string]*entity.GameObject // map of ALL objects in the game
	commands chan *ClientCommand // Inbound messages from the clients.
	register chan *Client // Register requests from the clients.
	unregister chan *Client // Unregister requests from clients.
}

func (e Engine) Floors() []*utils.Quadtree {
	return e.floors
}

func (e Engine) GameObjects() map[string]*entity.GameObject {
	return e.gameObjects
}

func (e Engine) Players() map[int]*entity.Player {
	return e.players
}

func (e Engine) SendResponse(responseType string, responseData map[string]interface{}, player *entity.Player) {
	resp := entity.EngineResponse{
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
		engine.UnregisterClient(e, player.Client)
	}
}

// send updates to all players who can see it
func (e Engine) SendResponseToVisionAreas(gameObj *entity.GameObject, responseType string, responseData map[string]interface{}) {
	intersectingObjects := e.Floors()[gameObj.Floor].RetrieveIntersections(utils.Bounds{
		X:      gameObj.X,
		Y:      gameObj.Y,
		Width:  gameObj.Width,
		Height: gameObj.Height,
	})
	resp := entity.EngineResponse{
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
						engine.UnregisterClient(e, player.Client)
					}
				}
			}
		}
	}
}

// send new state of the game object to all players who can see it
func (e Engine) SendGameObjectUpdate(gameObj *entity.GameObject, updateType string) {
	e.SendResponseToVisionAreas(gameObj, updateType, map[string]interface{}{
		"object": gameObj,
	})
	storage.GetClient().Updates <- gameObj
}

// used to send errors and other system response info
func (e Engine) SendSystemMessage(message string, player *entity.Player) {
	e.SendResponse("add_message", map[string]interface{}{
		"type": "system",
		"message": message,
	}, player)
}

func NewEngine() *Engine {
	return &Engine{
		tickTime:    0,
		players:     make(map[int]*entity.Player),
		gameObjects: make(map[string]*entity.GameObject),
		floors:      make([]*utils.Quadtree, FloorCount),
		commands:    make(chan *ClientCommand),
		register:    make(chan *Client),
		unregister:  make(chan *Client),
	}
}

func (e *Engine) Init() {
	// Start routine to process game objects updates and save them in game storage
	go storage.GetClient().Run()

	e.floors[0] = &utils.Quadtree{
		Bounds: utils.Bounds{
			X:      0,
			Y:      0,
			Width:  FloorSize,
			Height: FloorSize,
		},
		MaxObjects: 30,
		MaxLevels:  10,
		Level:      0,
		Objects:    make([]utils.IBounds, 0),
		Nodes:      make([]utils.Quadtree, 0),
	}
	engine.LoadGameObjects(e, FloorSize)
	e.tickTime = utils.MakeTimestamp()
}

// Main engine loop
func (e *Engine) Run() {
	e.Init()
	for {
		select {
		case client := <-e.register:
			engine.RegisterClient(e, client)
		case client := <-e.unregister:
			engine.UnregisterClient(e, client)
		case cmd := <-e.commands:
			engine.ProcessCommand(e, cmd.characterId, cmd.command)
		default:
			// Run world once in TickSize
			newTickTime := utils.MakeTimestamp()
			if newTickTime - e.tickTime >= TickSize {
				engine.MovePlayers(e, newTickTime - e.tickTime)
				e.tickTime = newTickTime
			}
		}
	}
}
