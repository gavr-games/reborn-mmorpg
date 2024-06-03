package door_object_test

import (
	"testing"

	"github.com/gavr-games/reborn-mmorpg/pkg/game_test"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test/factories"
	"github.com/stretchr/testify/assert"
)

func TestClose(t *testing.T) {
	testFunction = callClose
	e := game_test.GetEngine()
	gameObjectFactory := factories.NewGameObjectFactory()
	charGameObj := gameObjectFactory.CreateCharGameObject(game_test.GetEngine())
	player = gameObjectFactory.CreatePlayer(e, charGameObj)
	doorObject = gameObjectFactory.CreateObjectKeyXYArea(e, DoorObjectKey, charGameObj.X() + AwayCoord, charGameObj.Y(), charGameObj.GameAreaId())

	t.Run("Character needs to be closer", testNotClose)

	charGameObj.SetX(doorObject.X())
	charGameObj.SetY(doorObject.Y())
	t.Run("Door closed", testSuccess)

	t.Run("Door closed", func(t *testing.T) {
		state := doorObject.GetProperty("state")
		collidable := doorObject.GetProperty("collidable")
		assert.Equal(t, "closed", state)
		assert.True(t, collidable.(bool))
	})
}
