package game

import (
	"log"

	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine"
)

const FloorSize = 1000.0
const FloorCount = 4

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
		gameObjectsId: 0,
		gameObjects: make(map[int]*entity.GameObject),
		floors:      make([]*utils.Quadtree, FloorCount),
		commands:    make(chan *ClientCommand),
		register:    make(chan *Client),
		unregister:  make(chan *Client),
	}
}

func (e *Engine) Run() {
	engine.LoadGameObjects(e.floors, e.gameObjects, &e.gameObjectsId, FloorSize)
	//log.Println("Objects count %v", e.gameObjects[1])
	for {
		/*select {
		case client := <-e.register:
			e.clients[client] = true
		case client := <-e.unregister:
			if _, ok := e.clients[client]; ok {
				delete(e.clients, client)
				close(client.send)
			}
		case cmd := <-e.commands:
			for client := range e.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(e.clients, client)
				}
			}
		}*/
	}
}
