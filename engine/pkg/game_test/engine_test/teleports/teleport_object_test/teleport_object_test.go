package teleport_object_test

import (
	"os"
	"testing"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test"
	"github.com/stretchr/testify/assert"
)

const (
	CloseCoord = 0.5
	AwayCoord = 30.0
	TeleportObjectKey = "teleport/town_gate"
	TeleportToX = 10.1
	TeleportToY = 10.2
)

var charGameObj entity.IGameObject
var teleportObject entity.IGameObject
var testFunction func() (bool, error)

func TestMain(m *testing.M) {
	game_test.Setup()
	os.Exit(m.Run())
}

func testAreaNotFound(t *testing.T) {
	res, err := testFunction()
	assert.False(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, "area not found", err.Error())
}

func testNotClose(t *testing.T) {
	res, err := testFunction()
	assert.False(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, "too far from teleport", err.Error())
}

func testSuccess(t *testing.T) {
	res, err := testFunction()
	assert.True(t, res)
	assert.Nil(t, err)
}

func callTeleportTo() (bool, error) {
	return teleportObject.(entity.ITeleportObject).TeleportTo(game_test.GetEngine(), charGameObj)
}
