package game

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/puzpuzpuz/xsync/v3"
	uuid "github.com/satori/go.uuid"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/constants"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/armors/armor_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/bonfires/bonfire_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/characters"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/characters/character_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/claims/claim_obelisk_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/containers/backpack_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/containers/bag_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/containers/chest_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/delayed_actions"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/doors/door_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/dragons/dragon_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/effects"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/fishing/fishing_rod_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/hatcheries/hatchery_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mobs"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mobs/mob_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/npcs/npc_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/plants/plant_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/potions/potion_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/resources/resource_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/rocks/rock_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/shovels/shovel_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/teleports/teleport_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/tools/tool_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/trees/tree_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/walls/wall_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/weapons/weapon_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/vision_area_updater"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

// Engine runs the game
type Engine struct {
	tickTime       int64                                        //last tick time in milliseconds
	gameAreas      *xsync.MapOf[string, *entity.GameArea]       // map of ALL game areas, underground, etc
	players        *xsync.MapOf[int, *entity.Player]            // map of all players
	gameObjects    *xsync.MapOf[string, entity.IGameObject]     // map of ALL objects in the game
	delayedActions *xsync.MapOf[string, entity.IGameObject]     // map of objects with delayed actions
	mobs           *xsync.MapOf[string, entity.IMobObject]      // map of ALL mobs in the game
	effects        *xsync.MapOf[string, map[string]interface{}] // all active effects in the game
	commands       chan *ClientCommand                          // Inbound messages from the clients.
	register       chan *Client                                 // Register requests from the clients.
	unregister     chan *Client                                 // Unregister requests from clients.
	tasks          chan entity.Task                             // Messages from other routines to execute code in main engine loop
	testingMode    bool		                                     // tasks are performed immediately, not sent to channel
}

func (e *Engine) GameAreas() *xsync.MapOf[string, *entity.GameArea] {
	return e.gameAreas
}

func (e *Engine) GameObjects() *xsync.MapOf[string, entity.IGameObject] {
	return e.gameObjects
}

func (e *Engine) DelayedActions() *xsync.MapOf[string, entity.IGameObject] {
	return e.delayedActions
}

func (e *Engine) Mobs() *xsync.MapOf[string, entity.IMobObject] {
	return e.mobs
}

func (e *Engine) Players() *xsync.MapOf[int, *entity.Player] {
	return e.players
}

func (e *Engine) Effects() *xsync.MapOf[string, map[string]interface{}] {
	return e.effects
}

func (e *Engine) CurrentTickTime() int64 {
	return e.tickTime
}

// This func is used to send code to main engine loop from other goroutines
// This is required for critical object updates to be performed sequentially by engine
func (e *Engine) PerformTask(f func()) {
	if e.testingMode {
		f()
	} else {
		e.tasks <- f
	}
}

func (e *Engine) EnableTestingMode() {
	e.testingMode = true
}

func (e *Engine) GetGameAreaByKey(key string) *entity.GameArea {
	var gameArea *entity.GameArea
	e.GameAreas().Range(func(id string, ga *entity.GameArea) bool {
		if ga.Key() == key {
			gameArea = ga
			return false // stop iteration
		}
		return true
	})
	return gameArea
}

// Sends an update named responseType with parameters responseData to specific player (ONLY ONE).
func (e *Engine) SendResponse(responseType string, responseData map[string]interface{}, player *entity.Player) {
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
func (e *Engine) SendResponseToVisionAreas(gameObj entity.IGameObject, responseType string, responseData map[string]interface{}) {
	vision_area_updater.GetUpdater(e).Updates <- &vision_area_updater.VisionAreaUpdate{
		GameObj:      gameObj,
		ResponseType: responseType,
		ResponseData: responseData,
	}
}

// Send new update of the gameObj to all players who can see it
// IMPORTANT: this function also updates/delets gameObj in storage
func (e *Engine) SendGameObjectUpdate(gameObj entity.IGameObject, updateType string) {
	clone := gameObj.Clone() // clone is required to prevent access to objects map from different routines
	e.SendResponseToVisionAreas(gameObj, updateType, map[string]interface{}{
		"object": clone,
	})
	if updateType == "remove_object" {
		storage.GetClient().Deletes <- clone.Id()
	} else {
		storage.GetClient().Updates <- clone
	}
}

// Sends errors and other system response messages to specific player
func (e *Engine) SendSystemMessage(message string, player *entity.Player) {
	e.SendResponse("add_message", map[string]interface{}{
		"type":    "system",
		"message": message,
	}, player)
}

// Creates specific struct depending on object type and kind
// For example TreeObject for tree, RockObject for rock, etc.
func (e *Engine) CreateGameObjectStruct(gameObj entity.IGameObject) entity.IGameObject {
	switch gameObj.Type() {
	case "armor":
		return armor_object.NewArmorObject(gameObj)
	case "axe", "frying_pan", "hammer", "knife", "needle", "pickaxe", "saw":
		return tool_object.NewToolObject(gameObj)
	case "bonfire":
		return bonfire_object.NewBonfireObject(gameObj)
	case "claim":
		if gameObj.Kind() == "claim_obelisk" {
			return &claim_obelisk_object.ClaimObeliskObject{GameObject: *gameObj.(*entity.GameObject)}
		}
	case "container":
		if strings.Contains(gameObj.Kind(), "bag") {
			return bag_object.NewBagObject(gameObj)
		} else if strings.Contains(gameObj.Kind(), "chest") {
			return chest_object.NewChestObject(gameObj)
		} else {
			return backpack_object.NewBackpackObject(gameObj)
		}
	case "door":
		return door_object.NewDoorObject(gameObj)
	case "fishing_rod":
		return fishing_rod_object.NewFishingRodObject(gameObj)
	case "hatchery":
		return hatchery_object.NewHatcheryObject(gameObj)
	case "npc":
		return &npc_object.NpcObject{GameObject: *gameObj.(*entity.GameObject)}
	case "melee_weapon":
		return weapon_object.NewWeaponObject(gameObj)
	case "mob":
		if strings.Contains(gameObj.Kind(), "_dragon") {
			return dragon_object.NewDragonObject(e, gameObj)
		} else {
			return mob_object.NewMobObject(e, gameObj)
		}
	case "potion":
		return potion_object.NewPotionObject(gameObj)
	case "plant":
		return plant_object.NewPlantObject(gameObj)
	case "player":
		if gameObj.Kind() == "player" {
			return character_object.NewCharacterObject(gameObj)
		}
	case "resource":
		return resource_object.NewResourceObject(gameObj)
	case "rock":
		return &rock_object.RockObject{GameObject: *gameObj.(*entity.GameObject)}
	case "shovel":
		return shovel_object.NewShovelObject(gameObj)
	case "teleport":
		return teleport_object.NewTeleportObject(gameObj)
	case "tree":
		return &tree_object.TreeObject{GameObject: *gameObj.(*entity.GameObject)}
	case "wall", "well", "equipment", "sit":
		return wall_object.NewWallObject(gameObj)
	default:
		return gameObj
	}
	return gameObj
}

// Creates new GameObject and returns it
func (e *Engine) CreateGameObject(objPath string, x, y, rotation float64, gameAreaId string, additionalProps map[string]interface{}) entity.IGameObject {
	gameObj, err := game_objects.CreateFromTemplate(e, objPath, x, y, rotation)
	if err != nil {
		//TODO: handle error
	}
	if additionalProps != nil {
		for k, v := range additionalProps {
			gameObj.SetProperty(k, v)
		}
	}

	// Some templates might have actions to be created with the object
	if currentAction := gameObj.GetProperty("current_action"); currentAction != nil {
		actionParams := currentAction.(map[string]interface{})["params"].(map[string]interface{})
		actionParams["game_object_id"] = gameObj.Id()
		timeLeft := currentAction.(map[string]interface{})["time_left"].(float64)
		funcName := currentAction.(map[string]interface{})["func_name"].(string)
		delayedAction := entity.NewDelayedAction(
			funcName,
			actionParams,
			timeLeft,
			entity.DelayedActionReady,
		)
		gameObj.SetCurrentAction(delayedAction)
		gameObj.SetProperty("current_action", nil)
		e.DelayedActions().Store(gameObj.Id(), gameObj)
	}

	// Init effects
	for effectId, effect := range gameObj.Effects() {
		newEffectId := uuid.NewV4().String()
		effectMap := utils.CopyMap(effect.(map[string]interface{}))
		effectMap["id"] = effectId
		effectMap["target_id"] = gameObj.Id()
		e.Effects().Store(newEffectId, effectMap)
		gameObj.RemoveEffect(effectId)
		gameObj.SetEffect(newEffectId, effect)
	}

	gameObj.SetGameAreaId(gameAreaId)
	if gameAreaId != "" {
		if gameArea, gameAreaOk := e.GameAreas().Load(gameAreaId); gameAreaOk {
			gameArea.Insert(gameObj)
		}
	}

	e.GameObjects().Store(gameObj.Id(), gameObj)

	if gameObj.Kind() != "player_vision_area" {
		storage.GetClient().Updates <- gameObj.Clone()
	}

	if gameObj.Type() == "mob" {
		e.Mobs().Store(gameObj.Id(), gameObj.(entity.IMobObject))
	}

	return gameObj
}

func (e *Engine) RemoveGameObject(gameObj entity.IGameObject) {
	if gameObj.Type() == "mob" {
		e.Mobs().Delete(gameObj.Id())
	}
	e.DelayedActions().Delete(gameObj.Id())
	e.GameObjects().Delete(gameObj.Id())
	for effectId := range gameObj.Effects() {
		e.Effects().Delete(effectId)
	}
	if gameArea, gaOk := e.GameAreas().Load(gameObj.GameAreaId()); gaOk {
		gameObjId := gameObj.Id()
		gameArea.FilteredRemove(gameObj, func(b utils.IBounds) bool {
			return gameObjId == b.(entity.IGameObject).Id()
		})
	}
	e.SendGameObjectUpdate(gameObj, "remove_object")
}

func NewEngine() *Engine {
	return &Engine{
		tickTime:       0,
		players:        xsync.NewMapOf[int, *entity.Player](),
		gameObjects:    xsync.NewMapOf[string, entity.IGameObject](),
		delayedActions: xsync.NewMapOf[string, entity.IGameObject](),
		mobs:           xsync.NewMapOf[string, entity.IMobObject](),
		effects:        xsync.NewMapOf[string, map[string]interface{}](),
		gameAreas:      xsync.NewMapOf[string, *entity.GameArea](),
		commands:       make(chan *ClientCommand),
		register:       make(chan *Client),
		unregister:     make(chan *Client),
		tasks:          make(chan entity.Task),
		testingMode:    false,
	}
}

func (e *Engine) Init(skipWorldGeneration bool) {
	// Start routines to process game objects updates and save them in game storage
	go storage.GetClient().Run()
	// Start routine, which updates players about changes in their vision area
	vision_area_updater.GetUpdater(e).Run()

	e.tickTime = utils.MakeTimestamp()

	if !skipWorldGeneration {
		engine.LoadGameObjects(e) // Generate new worlds or read it from storage
	}
}

// Main engine loop
func (e *Engine) Run() {
	e.Init(false)
	for {
		select {
		case f := <- e.tasks: // we execute code from other routines, which has to be executed in main engine loop
			f()
		case client := <-e.register:
			engine.RegisterClient(e, client)
		case client := <-e.unregister:
			engine.UnregisterClient(e, client)
		case cmd := <-e.commands:
			// TODO: refactor to "go engine.ProcessCommand(e, cmd.characterId, cmd.command)"
			// it requires to figure out how to deal with slices and maps in game objects properties
			// there could be cases when 2 routines override a slice. 2 players put item at the same position in container.
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
