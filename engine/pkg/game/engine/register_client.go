package engine

import (
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
		VisibleObjects: make(map[string]bool),
	}
	e.Players()[player.Id] = player
	additionalProps := make(map[string]interface{})
	additionalProps["player_id"] = player.Id
	gameObj := CreateGameObject("player", InitialPlayerX, InitialPlayerY, additionalProps)
	gameObj.Floor = 0
	e.GameObjects()[gameObj.Id] = gameObj
	e.Floors()[gameObj.Floor].Insert(gameObj)
	player.CharacterGameObjectId = gameObj.Id
	CreatePlayerItems(e, player)
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

func CreatePlayerItems(e IEngine, player *entity.Player) {
	charGameObj := e.GameObjects()[player.CharacterGameObjectId]
	// Backpack
	additionalProps := make(map[string]interface{})
	additionalProps["owner_id"] = charGameObj.Id
	initialBackpack := CreateGameObject("backpack", charGameObj.X, charGameObj.Y, additionalProps)
	charGameObj.Properties["slots"].(map[string]interface{})["back"] = initialBackpack.Id
	initialBackpack.Floor = 0
	e.GameObjects()[initialBackpack.Id] = initialBackpack
	e.Floors()[initialBackpack.Floor].Insert(initialBackpack) // TODO: should we insert items in bags? Should we insert bag itself?
	// Axe
	initialAxe := CreateGameObject("axe", charGameObj.X, charGameObj.Y, nil)
	PutToContainer(e, player, initialBackpack.Id, initialAxe.Id, -1)
	initialAxe.Floor = 0
	e.GameObjects()[initialAxe.Id] = initialAxe
	e.Floors()[initialAxe.Floor].Insert(initialAxe) // TODO: should we insert items in bags?
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
			player.VisibleObjects = make(map[string]bool)
		}
		player.Client = client
	} else {
		player = CreatePlayer(e, client)
		CreatePlayerVisionArea(e, player)
	}
	if player, ok := e.Players()[client.GetCharacter().Id]; ok {
		visibleObjects := GetPlayerVisibleObjects(e, player)
		for _, val := range visibleObjects {
			player.VisibleObjects[val.(*entity.GameObject).Id] = true
		}
		//Send json with VisibleObjects from vision area
		SendResponse(e, "init_game", map[string]interface{}{
			"visible_objects": visibleObjects,
		}, player)
		SendGameObjectUpdate(e, e.GameObjects()[player.CharacterGameObjectId], "add_object")
	}
}