package leveling_object_test

import (
	"testing"

	"github.com/gavr-games/reborn-mmorpg/pkg/game_test"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test/factories"
	"github.com/stretchr/testify/assert"
)

const (
	chopTreeExp = 5.0
)

func TestAddExperience(t *testing.T) {
	testFunction = callAddExperience
	gameObjectFactory := factories.NewGameObjectFactory()
	charGameObj = gameObjectFactory.CreateCharGameObject(game_test.GetEngine())
	
	action = "wrong/action"
	t.Run("Action not found", testWrongAction)

	action = "chop_tree"
	t.Run("Add experience", testSuccess)
	t.Run("Experience has changed", func(t *testing.T) {
		amount := charGameObj.GetProperty("experience")
		level  := charGameObj.GetProperty("level")
		assert.Equal(t, chopTreeExp, amount)
		assert.Equal(t, 0.0, level)
	})

	action = "hatch/fire_dragon"
	t.Run("Add experience", testSuccess)
	t.Run("Add experience", testSuccess)
	t.Run("Level has changed", func(t *testing.T) {
		amount := charGameObj.GetProperty("experience")
		level  := charGameObj.GetProperty("level")
		assert.Equal(t, chopTreeExp, amount)
		assert.Equal(t, 1.0, level)
	})
}
