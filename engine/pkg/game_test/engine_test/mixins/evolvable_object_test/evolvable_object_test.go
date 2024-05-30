package evolvable_object_test

import (
	"os"
	"testing"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test"
	"github.com/stretchr/testify/assert"
)

const (
	AwayCoord  = 30.0
	EvolvableObjectKey = "mob/baby_fire_dragon"
	EvolveToKind = "fire_dragon"
)

var player *entity.Player
var evolvableObject entity.IGameObject
var testFunction func() (bool, error)

func TestMain(m *testing.M) {
	game_test.Setup()
	os.Exit(m.Run())
}

func testNotClose(t *testing.T) {
	res, err := testFunction()
	assert.False(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, "Character needs to be closer", err.Error())
}

func testNotEnoughFood(t *testing.T) {
	res, err := testFunction()
	assert.False(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, "Object needs more food to evolve", err.Error())
}

func testAlive(t *testing.T) {
	res, err := testFunction()
	assert.False(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, "The evolvable object is dead", err.Error())
}

func testSuccess(t *testing.T) {
	res, err := testFunction()
	assert.True(t, res)
	assert.Nil(t, err)
}

func callEvolve() (bool, error) {
	return evolvableObject.(entity.IEvolvableObject).Evolve(game_test.GetEngine(), player)
}
