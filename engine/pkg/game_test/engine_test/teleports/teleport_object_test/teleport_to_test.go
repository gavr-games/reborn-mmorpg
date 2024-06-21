package teleport_object_test

import (
	"testing"

	"github.com/gavr-games/reborn-mmorpg/pkg/game_test"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test/factories"
	"github.com/stretchr/testify/assert"
)

func TestTeleportTo(t *testing.T) {
	testFunction = callTeleportTo
	e := game_test.GetEngine()
	gameObjectFactory := factories.NewGameObjectFactory()
	charGameObj = gameObjectFactory.CreateCharGameObject(game_test.GetEngine())
	gameObjectFactory.CreatePlayer(e, charGameObj)
	gameObjectFactory.CreateVisionArea(e, charGameObj)
	teleportObject = gameObjectFactory.CreateObjectKeyXYArea(e, TeleportObjectKey, charGameObj.X() + AwayCoord, charGameObj.Y(), charGameObj.GameAreaId())

	t.Run("Test not close", testNotClose)

	teleportObject.SetProperty("teleport_to", map[string]interface{}{
		"area": "wrongtest",
	})
	t.Run("Test not close", testAreaNotFound)

	teleportObject.SetX(charGameObj.X())
	teleportObject.SetProperty("teleport_to", map[string]interface{}{
		"area": "surface",
		"x": TeleportToX,
		"y": TeleportToY,
	})

	t.Run("Teleported", testSuccess)

	t.Run("Char coords has changed", func(t *testing.T) {
		assert.Equal(t, TeleportToX, charGameObj.X())
		assert.Equal(t, TeleportToY, charGameObj.Y())
	})

	teleportObject.SetX(charGameObj.X())
	teleportObject.SetY(charGameObj.Y())
	teleportObject.SetProperty("teleport_to", map[string]interface{}{
		"area": "surface",
		"random_coords": true,
	})

	t.Run("Randomly teleported", testSuccess)

	t.Run("Char coords has changed randomly", func(t *testing.T) {
		assert.NotEqual(t, TeleportToX, charGameObj.X())
		assert.NotEqual(t, TeleportToY, charGameObj.Y())
	})
}
