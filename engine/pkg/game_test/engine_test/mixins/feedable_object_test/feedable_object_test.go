package feedable_object_test

import (
	"os"
	"testing"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test"
	"github.com/stretchr/testify/assert"
)

const (
	AwayCoord  = 30.0
	FeedableObjectKey = "mob/baby_fire_dragon"
	EatableObjectKey = "resource/carrot"
	NotEatableObjectKey = "resource/stone"
)

var player *entity.Player
var feedableObject entity.IGameObject
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
	assert.Equal(t, "Character needs to be closer", err.Error())
}

func testNotEatable(t *testing.T) {
	res, err := testFunction()
	assert.False(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, "Wrong item to feed", err.Error())
}

func testNotInContainer(t *testing.T) {
	res, err := testFunction()
	assert.False(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, "Item is not in container", err.Error())
}

func testSuccess(t *testing.T) {
	res, err := testFunction()
	assert.True(t, res)
	assert.Nil(t, err)
}

func callFeed() (bool, error) {
	return feedableObject.(entity.IFeedableObject).Feed(game_test.GetEngine(), itemId, player)
}
