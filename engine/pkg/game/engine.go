package game

import (
	"encoding/json"
	"fmt"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/constants"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/plants/plant_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/characters"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/characters/character_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/claims/claim_obelisk_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/containers/backpack_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/containers/bag_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/delayed_actions"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/effects"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/hatcheries/hatchery_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mobs"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mobs/mob_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/npcs/npc_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/potions/potion_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/resources/resource_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/rocks/rock_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/shovels/shovel_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/tools/tool_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/trees/tree_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/walls/wall_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/weapons/weapon_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

// Engine runs the game
type Engine struct {
	tickTime    int64                             //last tick time in milliseconds
	floors      []*utils.Quadtree                 // slice of global game areas, underground, etc
	players     map[int]*entity.Player            // map of all players
	gameObjects map[string]entity.IGameObject     // map of ALL objects in the game
	mobs        map[string]entity.IMobObject      // map of ALL mobs in the game
	effects     map[string]map[string]interface{} // all active effects in the game
	commands    chan *ClientCommand               // Inbound messages from the clients.
	register    chan *Client                      // Register requests from the clients.
	unregister  chan *Client                      // Unregister requests from clients.
}

func (e Engine) Floors() []*utils.Quadtree {
	return e.floors
}

func (e Engine) GameObjects() map[string]entity.IGameObject {
	return e.gameObjects
}

func (e Engine) Mobs() map[string]entity.IMobObject {
	return e.mobs
}

func (e Engine) Players() map[int]*entity.Player {
	return e.players
}

func (e Engine) Effects() map[string]map[string]interface{} {
	return e.effects
}

func (e Engine) CurrentTickTime() int64 {
	return e.tickTime
}

// Sends an update named responseType with parameters responseData to specific player (ONLY ONE).
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
	if player.Client != nil {
		select {
		case player.Client.GetSendChannel() <- message:
		default:
			engine.UnregisterClient(e, player.Client)
		}
	}
}

// Sends an update named responseType with parameters responseData to all players,
// who can see the gameObj. In other words their vision areas collide with gameObj X,Y.
func (e Engine) SendResponseToVisionAreas(gameObj entity.IGameObject, responseType string, responseData map[string]interface{}) {
	intersectingObjects := e.Floors()[gameObj.Floor()].RetrieveIntersections(utils.Bounds{
		X:      gameObj.X(),
		Y:      gameObj.Y(),
		Width:  gameObj.Width(),
		Height: gameObj.Height(),
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
		if obj.(entity.IGameObject).Type() == "player" && obj.(entity.IGameObject).Kind() == "player_vision_area" {
			playerId := obj.(entity.IGameObject).Properties()["player_id"].(int)
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

// Send new update of the gameObj to all players who can see it
// IMPORTANT: this function also updates/delets gameObj in storage
func (e Engine) SendGameObjectUpdate(gameObj entity.IGameObject, updateType string) {
	clone := gameObj.Clone() // clone is required to prevent access to objects map from different routines
	e.SendResponseToVisionAreas(gameObj, updateType, map[string]interface{}{
		"object": gameObj,
	})
	if updateType == "remove_object" {
		storage.GetClient().Deletes <- clone.Id()
	} else {
		storage.GetClient().Updates <- clone
	}
}

// Sends errors and other system response messages to specific player
func (e Engine) SendSystemMessage(message string, player *entity.Player) {
	e.SendResponse("add_message", map[string]interface{}{
		"type":    "system",
		"message": message,
	}, player)
}

// Creates specific struct depending on object type and kind
// For example TreeObject for tree, RockObject for rock, etc.
func (e Engine) CreateGameObjectStruct(gameObj entity.IGameObject) entity.IGameObject {
	switch gameObj.Type() {
	case "tree":
		return &tree_object.TreeObject{*gameObj.(*entity.GameObject)}
	case "rock":
		return &rock_object.RockObject{*gameObj.(*entity.GameObject)}
	case "potion":
		return potion_object.NewPotionObject(gameObj)
	case "plant":
		return &plant_object.PlantObject{*gameObj.(*entity.GameObject)}
	case "hatchery":
		return hatchery_object.NewHatcheryObject(gameObj)
	case "wall":
		return wall_object.NewWallObject(gameObj)
	case "npc":
		return &npc_object.NpcObject{*gameObj.(*entity.GameObject)}
	case "resource":
		return resource_object.NewResourceObject(gameObj)
	case "melee_weapon":
		return weapon_object.NewWeaponObject(gameObj)
	case "hammer", "knife", "pickaxe", "axe", "needle", "fishing_rod":
		return tool_object.NewToolObject(gameObj)
	case "shovel":
		return shovel_object.NewShovelObject(gameObj)
	case "mob":
		return mob_object.NewMobObject(e, gameObj)
	case "player":
		if gameObj.Kind() == "player" {
			return character_object.NewCharacterObject(gameObj)
		}
	case "container":
		if gameObj.Kind() == "backpack" {
			return backpack_object.NewBackpackObject(gameObj)
		} else {
			return bag_object.NewBagObject(gameObj)
		}
	case "claim":
		if gameObj.Kind() == "claim_obelisk" {
			return &claim_obelisk_object.ClaimObeliskObject{*gameObj.(*entity.GameObject)}
		}
	default:
		return gameObj
	}
	return gameObj
}

// Creates new GameObject and returns it
func (e Engine) CreateGameObject(objPath string, x float64, y float64, rotation float64, floor int, additionalProps map[string]interface{}) entity.IGameObject {
	gameObj, err := game_objects.CreateFromTemplate(e, objPath, x, y, rotation)
	if err != nil {
		//TODO: handle error
	}
	if additionalProps != nil {
		for k, v := range additionalProps {
			gameObj.Properties()[k] = v
		}
	}

	gameObj.SetFloor(floor)
	if floor != -1 {
		e.Floors()[gameObj.Floor()].Insert(gameObj)
	}

	e.GameObjects()[gameObj.Id()] = gameObj

	if gameObj.Kind() != "player_vision_area" {
		storage.GetClient().Updates <- gameObj.Clone()
	}

	if gameObj.Properties()["type"].(string) == "mob" {
		e.Mobs()[gameObj.Id()] = gameObj.(entity.IMobObject)
	}

	return gameObj
}

func NewEngine() *Engine {
	return &Engine{
		tickTime:    0,
		players:     make(map[int]*entity.Player),
		gameObjects: make(map[string]entity.IGameObject),
		mobs:        make(map[string]entity.IMobObject),
		effects:     make(map[string]map[string]interface{}),
		floors:      make([]*utils.Quadtree, constants.FloorCount),
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
			Width:  constants.FloorSize,
			Height: constants.FloorSize,
		},
		MaxObjects: 30,
		MaxLevels:  10,
		Level:      0,
		Objects:    make([]utils.IBounds, 0),
		Nodes:      make([]utils.Quadtree, 0),
	}
	engine.LoadGameObjects(e) // Generate new worlds or read it from storage
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
			tickDelta := newTickTime-e.tickTime
			if tickDelta >= constants.TickSize {
				characters.Update(e, tickDelta)
				mobs.Update(e, tickDelta, newTickTime)
				effects.Update(e, tickDelta)
				delayed_actions.UpdateAll(e, tickDelta)
				e.tickTime = newTickTime
			}
		}
	}
}
