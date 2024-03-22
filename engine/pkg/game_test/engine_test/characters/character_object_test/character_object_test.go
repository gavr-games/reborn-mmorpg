package character_object_test

import (
	"os"
	"testing"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test"
	"github.com/stretchr/testify/assert"
)

const (
	totalGoldAmount     = 20.0
	substractGoldAmount = 10.0
)

var charGameObj entity.IGameObject
var amount float64
var testFunction func() (bool, error)

func TestMain(m *testing.M) {
	game_test.Setup()
	os.Exit(m.Run())
}

func testPlayerDoesNotExist(t *testing.T) {
	res, err := testFunction()
	assert.False(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, "Player does not exist", err.Error())
}

func testPlayerDoesNotHaveContainer(t *testing.T) {
	res, err := testFunction()
	assert.False(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, "Player does not have container", err.Error())
}

func testPlayerDoesNotHaveRequiredResources(t *testing.T) {
	res, err := testFunction()
	assert.False(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, "Player does not have required resources", err.Error())
}

func testSuccess(t *testing.T) {
	res, err := testFunction()
	assert.True(t, res)
	assert.Nil(t, err)
}

func callSubstractGold() (bool, error) {
	return charGameObj.(entity.ICharacterObject).SubstractGold(game_test.GetEngine(), amount)
}
