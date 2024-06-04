package burning_object_test

import (
	"testing"

	"github.com/gavr-games/reborn-mmorpg/pkg/game_test"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test/factories"
	"github.com/stretchr/testify/assert"
)

func TestBurn(t *testing.T) {
	testFunction = callBurn
	e := game_test.GetEngine()
	gameObjectFactory := factories.NewGameObjectFactory()
	charGameObj := gameObjectFactory.CreateCharGameObject(game_test.GetEngine())
	player = gameObjectFactory.CreatePlayer(e, charGameObj)
	burningObject = gameObjectFactory.CreateObjectKeyXYArea(e, BurningObjectKey, charGameObj.X() + AwayCoord, charGameObj.Y(), charGameObj.GameAreaId())

	t.Run("Character needs to be closer", testNotClose)

	charGameObj.SetX(burningObject.X())
	charGameObj.SetY(burningObject.Y())
	t.Run("No fuel", testNoFuel)

	burningObject.SetProperty("fuel", burningObject.GetProperty("max_fuel"))
	burningObject.SetProperty("state", "burning")
	t.Run("Already burning", testAlreadyBurning)

	burningObject.SetProperty("state", "extinguished")
	t.Run("Start burning", testSuccess)

	t.Run("State has changed", func(t *testing.T) {
		state := burningObject.GetProperty("state")
		effects := burningObject.Effects()
		assert.Equal(t, "burning", state)
		assert.Equal(t, 1, len(effects))
	})
}
