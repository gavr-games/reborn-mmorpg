package engine

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/constants"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/serializers"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

func CreatePlayer(e entity.IEngine, client entity.IClient) *entity.Player {
	character := client.GetCharacter()
	player := &entity.Player{
		Id: character.Id,
		CharacterGameObjectId: "",
		VisionAreaGameObjectId: "",
		Client: client,
	}
	e.Players().Store(player.Id, player)
	additionalProps := make(map[string]interface{})
	additionalProps["player_id"] = player.Id
	additionalProps["name"] = character.Name
	surfaceArea := e.GetGameAreaByKey("surface")
	gameObj := e.CreateGameObject("player/player", constants.InitialPlayerX, constants.InitialPlayerY, 0.0, surfaceArea.Id(), additionalProps)
	player.CharacterGameObjectId = gameObj.Id()
	CreatePlayerItems(e, player)
	return player
}

func CreatePlayerVisionArea(e entity.IEngine, player *entity.Player) {
	if charGameObj, ok := e.GameObjects().Load(player.CharacterGameObjectId); ok {
		additionalProps := make(map[string]interface{})
		additionalProps["player_id"] = player.Id
		gameObj := e.CreateGameObject("player/player_vision_area", charGameObj.(entity.ICharacterObject).GetVisionAreaX(), charGameObj.(entity.ICharacterObject).GetVisionAreaY(), 0.0, charGameObj.GameAreaId(), additionalProps)
		player.VisionAreaGameObjectId = gameObj.Id()
	}
}

func CreatePlayerItems(e entity.IEngine, player *entity.Player) {
	if charGameObj, ok := e.GameObjects().Load(player.CharacterGameObjectId); ok {
		// Backpack
		additionalProps := make(map[string]interface{})
		additionalProps["owner_id"] = charGameObj.Id()
		initialBackpack := e.CreateGameObject("container/backpack", charGameObj.X(), charGameObj.Y(), 0.0, "", additionalProps)
		slots := charGameObj.GetProperty("slots").(map[string]interface{})
		slots["back"] = initialBackpack.Id()
		charGameObj.SetProperty("slots", slots)
		// Axe
		initialAxe := e.CreateGameObject("axe/axe", charGameObj.X(), charGameObj.Y(), 0.0, "", nil)
		initialBackpack.(entity.IContainerObject).Put(e, player, initialAxe.Id(), -1)
		// Pickaxe
		initialPickaxe := e.CreateGameObject("pickaxe/pickaxe", charGameObj.X(), charGameObj.Y(), 0.0, "", nil)
		initialBackpack.(entity.IContainerObject).Put(e, player, initialPickaxe.Id(), -1)
	}
}

// Process when new player logs into the game
func RegisterClient(e entity.IEngine, client entity.IClient) {
	// lookup if engine has already created player object
	if player, ok := e.Players().Load(client.GetCharacter().Id); ok {
		if player.Client != nil {
			// close previous socket connection for this player
			close(player.Client.GetSendChannel())
		} else {
			CreatePlayerVisionArea(e, player)
			if charGameObj, charOk := e.GameObjects().Load(player.CharacterGameObjectId); charOk {
				charGameObj.SetProperty("visible", true)
			}
		}
		player.Client = client
	} else {
		player = CreatePlayer(e, client)
		CreatePlayerVisionArea(e, player)
	}
	if player, ok := e.Players().Load(client.GetCharacter().Id); ok {
		if visionArea, visionAreaOk := e.GameObjects().Load(player.VisionAreaGameObjectId); visionAreaOk {
			go initGame(e, player, visionArea)
			if charGameObj, charOk := e.GameObjects().Load(player.CharacterGameObjectId); charOk {
				// Send character obj to another players
				charClone := charGameObj.Clone()
				charClone.SetProperties(serializers.GetInfo(e, charClone))
				e.SendResponseToVisionAreas(charGameObj, "add_object", map[string]interface{}{
					"object": charClone,
				})
				storage.GetClient().Updates <- charGameObj.Clone()
				// Show lifted object
				if liftedObjectId := charGameObj.GetProperty("lifted_object_id"); liftedObjectId != nil {
					if liftedObj, liftedObjOk := e.GameObjects().Load(liftedObjectId.(string)); liftedObjOk {
						if liftedObj != nil {
							liftedObj.SetProperty("visible", true)
							e.SendGameObjectUpdate(liftedObj, "add_object")
						}
					}
				}
			}
		}
	}
}

func initGame(e entity.IEngine, player *entity.Player, visionArea entity.IGameObject) {
	visibleObjects := game_objects.GetVisibleObjects(e, visionArea.GameAreaId(), visionArea.HitBox())
	for key, val := range visibleObjects {
		clone := val.(entity.IGameObject).Clone()
		// This is required to send target info on first character object rendering
		if val.(entity.IGameObject).Id() == player.CharacterGameObjectId {
			clone.SetProperties(serializers.GetInfo(e, clone))
		}
		visibleObjects[key] = clone
	}
	// Send json with VisibleObjects from vision area
	e.SendResponse("init_game", map[string]interface{}{
		"visible_objects": visibleObjects,
	}, player)
}
