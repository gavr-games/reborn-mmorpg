package burning_object_test

import (
	"os"
	"testing"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test"
	"github.com/stretchr/testify/assert"
)

const (
	AwayCoord  = 30.0
	BurningObjectKey = "bonfire/bonfire"
	FuelObjectKey = "resource/log"
	NotFuelObjectKey = "resource/carrot"
	FuelValue = 10000.0
)

var player *entity.Player
var burningObject entity.IGameObject
var itemId string
var testFunction func() (bool, error)

func TestMain(m *testing.M) {
	game_test.Setup()
	os.Exit(m.Run())
}

func testNotClose(t *testing.T) {
	res, err := testFunction()
	assert.False(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, "character needs to be closer", err.Error())
}

func testNotFuel(t *testing.T) {
	res, err := testFunction()
	assert.False(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, "this fuel is not allowed", err.Error())
}

func testFuelIsFull(t *testing.T) {
	res, err := testFunction()
	assert.False(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, "no space for more fuel", err.Error())
}

func testNoFuel(t *testing.T) {
	res, err := testFunction()
	assert.False(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, "no fuel", err.Error())
}

func testAlreadyBurning(t *testing.T) {
	res, err := testFunction()
	assert.False(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, "already burning", err.Error())
}

func testNotBurning(t *testing.T) {
	res, err := testFunction()
	assert.False(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, "not burning", err.Error())
}

func testNotInContainer(t *testing.T) {
	res, err := testFunction()
	assert.False(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, "item is not in container", err.Error())
}

func testSuccess(t *testing.T) {
	res, err := testFunction()
	assert.True(t, res)
	assert.Nil(t, err)
}

func callAddFuel() (bool, error) {
	return burningObject.(entity.IBurningObject).AddFuel(game_test.GetEngine(), itemId, player)
}

func callBurn() (bool, error) {
	return burningObject.(entity.IBurningObject).Burn(game_test.GetEngine(), player)
}

func callExtinguish() (bool, error) {
	return burningObject.(entity.IBurningObject).Extinguish(game_test.GetEngine(), player)
}
