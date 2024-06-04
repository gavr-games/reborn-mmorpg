package fishing_rod_object_test

import (
	"os"
	"testing"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test"
	"github.com/stretchr/testify/assert"
)

const (
	CloseCoord = 0.5
	FishingRodObjectKey = "fishing_rod/wooden_fishing_rod"
)

var charGameObj entity.IGameObject
var fishingRodObject entity.IGameObject
var testFunction func() (bool, error)

func TestMain(m *testing.M) {
	game_test.Setup()
	os.Exit(m.Run())
}

func testNotEquipped(t *testing.T) {
	res, err := testFunction()
	assert.False(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, "fishing rod is not equipped", err.Error())
}

func testNotNearWater(t *testing.T) {
	res, err := testFunction()
	assert.False(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, "character is not near the water", err.Error())
}

func testSuccess(t *testing.T) {
	res, err := testFunction()
	assert.True(t, res)
	assert.Nil(t, err)
}

func callCatch() (bool, error) {
	return fishingRodObject.(entity.IFishingRodObject).Catch(game_test.GetEngine(), charGameObj)
}
