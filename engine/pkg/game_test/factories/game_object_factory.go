package factories

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

type GameObjectFactory struct{}

func NewGameObjectFactory() *GameObjectFactory {
	return &GameObjectFactory{}
}

const (
	playerObjKey   = "player/player"
	npcObjKey      = "npc/town_keeper"
	backpackObjKey = "container/backpack"
)

func (f *GameObjectFactory) CreateCharGameObject(e *game.Engine) entity.IGameObject {
	maxId := findMaxPlayerId(e)
	return e.CreateGameObject(playerObjKey, 0.0, 0.0, 0, 0, map[string]interface{}{"player_id": maxId + 1})
}

func (f *GameObjectFactory) CreatePlayer(e *game.Engine, charGameObj entity.IGameObject) *entity.Player {
	playerId := charGameObj.Properties()["player_id"].(int)
	player :=  &entity.Player{Id: playerId, CharacterGameObjectId: charGameObj.Id()}
	e.Players().Store(playerId, player)
	return player
}

func (f *GameObjectFactory) CreateNpcGameObject(e *game.Engine) entity.IGameObject {
	return e.CreateGameObject(npcObjKey, 0, 0, 0, 0, nil)
}

func (f *GameObjectFactory) CreateBackpackGameObject(e *game.Engine, charGameObj entity.IGameObject) entity.IGameObject {
	return e.CreateGameObject(backpackObjKey, charGameObj.X(), charGameObj.Y(), 0.0, -1, map[string]interface{}{"owner_id": charGameObj.Id()})
}

func (f *GameObjectFactory) CreateResourceGameObject(e *game.Engine, charGameObj entity.IGameObject, resourceKey string) entity.IGameObject {
	return e.CreateGameObject(resourceKey, charGameObj.X(), charGameObj.Y(), 0.0, charGameObj.Floor(), map[string]interface{}{
		"visible": false,
	})
}

func (f *GameObjectFactory) CreateStackableResourceGameObject(e *game.Engine, charGameObj entity.IGameObject, resourceKey string, amount float64) entity.IGameObject {
	return e.CreateGameObject(resourceKey, charGameObj.X(), charGameObj.Y(), 0.0, charGameObj.Floor(), map[string]interface{}{
		"visible": false,
		"amount":  amount,
	})
}

func findMaxPlayerId(e *game.Engine) int {
	maxId := 1
	e.Players().Range(func(playerId int, player *entity.Player) bool {
		if playerId > maxId {
			maxId = playerId
		}
		return true
	})

	return maxId
}
