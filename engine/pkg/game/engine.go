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

func NewEngine() *Engine {
	return &Engine{
		players:     make(map[int]*entity.Player),
		gameObjects: make(map[string]*entity.GameObject),
		floors:      make([]*utils.Quadtree, FloorCount),
		commands:    make(chan *ClientCommand),
		register:    make(chan *Client),
		unregister:  make(chan *Client),
	}
}

func (e *Engine) Init() {
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
}

// Main engine communication loop
func (e *Engine) Run() {
	for {
		select {
		case client := <-e.register:
			engine.RegisterClient(e, client)
		case client := <-e.unregister:
			engine.UnregisterClient(e, client)
		default:
		case cmd := <-e.commands:
			log.Println(cmd)
		}
	}
}

func (e *Engine) RunWorld() {
	e.Init()
	for {
		// move players
	}
}
