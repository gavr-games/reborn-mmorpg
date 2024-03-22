package claim_obelisk_object_test

import (
	"os"
	"testing"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test"
	"github.com/stretchr/testify/assert"
)

const (
	objectKind = "dragon_altar"
	objectKey  = "equipment/dragon_altar"
	awayCoord  = 30.0
	nearCoord  = 2.0
)

var charGameObj entity.IGameObject
var claimObeliskObj entity.IGameObject
var kind string
var testFunction func() (interface{}, error)
var res interface{}
var err error

func TestMain(m *testing.M) {
	game_test.Setup()
	os.Exit(m.Run())
}

func testClaimAreaDoesNotExist(t *testing.T) {
	res, err = testFunction()
	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, "Claim area does not exist", err.Error())
}

func testObjectNotFound(t *testing.T) {
	res, err = testFunction()
	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, "No object found on claim area", err.Error())
}

func testSuccess(t *testing.T) {
	res, err = testFunction()
	assert.NotNil(t, res)
	assert.Nil(t, err)
}

func callFindKindInArea() (interface{}, error) {
	return claimObeliskObj.(entity.IClaimObeliskObject).FindKindInArea(game_test.GetEngine(), kind)
}
