package burning_object_test

import (
	"testing"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test/factories"
	"github.com/stretchr/testify/assert"
)

func TestExtinguish(t *testing.T) {
	testFunction = callExtinguish
	e := game_test.GetEngine()
	gameObjectFactory := factories.NewGameObjectFactory()
	charGameObj := gameObjectFactory.CreateCharGameObject(game_test.GetEngine())
	player = gameObjectFactory.CreatePlayer(e, charGameObj)
	burningObject = gameObjectFactory.CreateObjectKeyXYArea(e, BurningObjectKey, charGameObj.X() + AwayCoord, charGameObj.Y(), charGameObj.GameAreaId())

	t.Run("Character needs to be closer", testNotClose)

	charGameObj.SetX(burningObject.X())
	charGameObj.SetY(burningObject.Y())
	t.Run("Not burning", testNotBurning)

	burningObject.SetProperty("fuel", burningObject.GetProperty("max_fuel"))
	burningObject.(entity.IBurningObject).Burn(e, player)
	t.Run("Extinguish fire", testSuccess)

	t.Run("State has changed", func(t *testing.T) {
		state := burningObject.GetProperty("state")
		effects := burningObject.Effects()
		assert.Equal(t, "extinguished", state)
		assert.Equal(t, 0, len(effects))
	})
}
