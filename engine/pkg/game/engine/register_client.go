package engine

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/constants"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/containers"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/serializers"
)

func CreatePlayer(e entity.IEngine, client entity.IClient) *entity.Player {
	player := &entity.Player{
		Id: client.GetCharacter().Id,
		CharacterGameObjectId: "",
		VisionAreaGameObjectId: "",
		Client: client,
		VisibleObjects: make(map[string]bool),
	}
	e.Players()[player.Id] = player
	additionalProps := make(map[string]interface{})
	additionalProps["player_id"] = player.Id
	gameObj := e.CreateGameObject("player/player", constants.InitialPlayerX, constants.InitialPlayerY, 0.0, 0, additionalProps)
	player.CharacterGameObjectId = gameObj.Id
	CreatePlayerItems(e, player)
	return player
}

func CreatePlayerVisionArea(e entity.IEngine, player *entity.Player) *entity.GameObject {
	charGameObj := e.GameObjects()[player.CharacterGameObjectId]
	additionalProps := make(map[string]interface{})
	additionalProps["player_id"] = player.Id
	gameObj := e.CreateGameObject("player/player_vision_area", charGameObj.X - game_objects.PlayerVisionArea / 2, charGameObj.Y - game_objects.PlayerVisionArea / 2, 0.0, 0, additionalProps)
	player.VisionAreaGameObjectId = gameObj.Id
	return gameObj
}

func CreatePlayerItems(e entity.IEngine, player *entity.Player) {
	charGameObj := e.GameObjects()[player.CharacterGameObjectId]
	// Backpack
	additionalProps := make(map[string]interface{})
	additionalProps["owner_id"] = charGameObj.Id
	initialBackpack := e.CreateGameObject("container/backpack", charGameObj.X, charGameObj.Y, 0.0, -1, additionalProps)
	charGameObj.Properties["slots"].(map[string]interface{})["back"] = initialBackpack.Id
	// Axe
	initialAxe := e.CreateGameObject("axe/axe", charGameObj.X, charGameObj.Y, 0.0, -1, nil)
	containers.Put(e, player, initialBackpack.Id, initialAxe.Id, -1)
	// Pickaxe
	initialPickaxe := e.CreateGameObject("pickaxe/pickaxe", charGameObj.X, charGameObj.Y, 0.0, -1, nil)
	containers.Put(e, player, initialBackpack.Id, initialPickaxe.Id, -1)
}

// Process when new player logs into the game
func RegisterClient(e entity.IEngine, client entity.IClient) {
	// lookup if engine has already created player object
	if player, ok := e.Players()[client.GetCharacter().Id]; ok {
		if player.Client != nil {
			// close previous socket connection for this player
			close(player.Client.GetSendChannel())
		} else {
			CreatePlayerVisionArea(e, player)
			e.GameObjects()[player.CharacterGameObjectId].Properties["visible"] = true
			player.VisibleObjects = make(map[string]bool)
		}
		player.Client = client
	} else {
		player = CreatePlayer(e, client)
		CreatePlayerVisionArea(e, player)
	}
	if player, ok := e.Players()[client.GetCharacter().Id]; ok {
		visibleObjects := GetPlayerVisibleObjects(e, player)
		for key, val := range visibleObjects {
			player.VisibleObjects[val.(*entity.GameObject).Id] = true
			// This is required to send target info on first character object rendering
			if val.(*entity.GameObject).Id == player.CharacterGameObjectId {
				clone := game_objects.Clone(val.(*entity.GameObject))
				clone.Properties = serializers.GetInfo(e.GameObjects(), val.(*entity.GameObject))
				visibleObjects[key] = clone
			}
		}
		// Send json with VisibleObjects from vision area
		e.SendResponse("init_game", map[string]interface{}{
			"visible_objects": visibleObjects,
		}, player)
		// Send character obj to another players
		e.SendGameObjectUpdate(e.GameObjects()[player.CharacterGameObjectId], "add_object")
	}
}