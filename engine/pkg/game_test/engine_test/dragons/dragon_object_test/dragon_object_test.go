package dragon_object_test

import (
	"os"
	"testing"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test"
	"github.com/stretchr/testify/assert"
)

const (
	altarKey  = "equipment/dragon_altar"
	dragonKey = "mob/fire_dragon"
	goldAmount = 100.0
	nearCoord  = 2.0
)

var charGameObj entity.IGameObject
var dragonObj entity.IGameObject
var claimObeliskObj entity.IGameObject
var testFunction func() (bool, error)

func TestMain(m *testing.M) {
	game_test.Setup()
	os.Exit(m.Run())
}

func testNotOwner(t *testing.T) {
	res, err := testFunction()
	assert.False(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, "Player is not the owner of this creature", err.Error())
}

func testAlreadyAlive(t *testing.T) {
	res, err := testFunction()
	assert.False(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, "The dragon is already alive", err.Error())
}

func testPlayerDoesNotHaveRequiredResources(t *testing.T) {
	res, err := testFunction()
	assert.False(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, "Player does not have required resources", err.Error())
}

func testClaimNotExists(t *testing.T) {
	res, err := testFunction()
	assert.False(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, "Claim does not exist", err.Error())
}

func testDragonAltarNotExists(t *testing.T) {
	res, err := testFunction()
	assert.False(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, "Dragon altar does not exist", err.Error())
}

func testSuccess(t *testing.T) {
	res, err := testFunction()
	assert.True(t, res)
	assert.Nil(t, err)
}

func callResurrect() (bool, error) {
	return dragonObj.(entity.IDragonObject).Resurrect(charGameObj)
}
