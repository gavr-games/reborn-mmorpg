package door_object_test

import (
	"os"
	"testing"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test"
	"github.com/stretchr/testify/assert"
)

const (
	AwayCoord  = 30.0
	DoorObjectKey = "door/wooden_door"
)

var player *entity.Player
var doorObject entity.IGameObject
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

func testSuccess(t *testing.T) {
	res, err := testFunction()
	assert.True(t, res)
	assert.Nil(t, err)
}

func callClose() (bool, error) {
	return doorObject.(entity.IDoorObject).Close(game_test.GetEngine(), player)
}

func callOpen() (bool, error) {
	return doorObject.(entity.IDoorObject).Open(game_test.GetEngine(), player)
}
