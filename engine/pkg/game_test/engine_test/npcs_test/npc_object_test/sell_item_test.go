package npc_object_test

import (
	"testing"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test/factories"
)

func TestSellItem(t *testing.T) {
	testFunction = callSellItem
	gameObjectFactory := factories.NewGameObjectFactory()
	charGameObj = gameObjectFactory.CreateCharGameObject(e)
	npcObj = gameObjectFactory.CreateNpcGameObject(e)

	t.Run("Player does not exist", testPlayerDoesNotExist)

	// Create a new player
	playerId := charGameObj.Properties()["player_id"].(int)
	e.Players()[playerId] = &entity.Player{Id: playerId, CharacterGameObjectId: charGameObj.Id()}
	player := e.Players()[playerId]

	t.Run("Player does not have container", testPlayerDoesNotHaveContainer)

	// Create a container for player
	slots := charGameObj.Properties()["slots"].(map[string]interface{})
	slots["back"] = gameObjectFactory.CreateBackpackGameObject(e, charGameObj).Id()

	// Place player far from NPC
	charGameObj.SetX(2.0)
	charGameObj.SetY(2.0)

	t.Run("Player need to be closer to NPC", testPlayerNeedToBeCloser)

	// Place player closer to NPC
	charGameObj.SetX(0)
	charGameObj.SetY(0)

	t.Run("Player does not have required resources", testPlayerDoesNotHaveRequiredResources)

	// Give player enough resources
	itemToSellObj := gameObjectFactory.CreateResourceGameObject(e, charGameObj, itemToSellObjKey)
	container := e.GameObjects()[slots["back"].(string)]
	container.(entity.IContainerObject).Put(e, player, itemToSellObj.Id(), -1)

	t.Run("Player sold item successfully", testSuccess)
}
