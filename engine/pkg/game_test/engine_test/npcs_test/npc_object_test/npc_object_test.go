package npc_object_test

import (
	"os"
	"testing"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test"
	"github.com/stretchr/testify/assert"
)

var npcObj entity.IGameObject
var charGameObj entity.IGameObject
var testFunction func() (bool, error)

const (
	amount           = 1.0
	sellingPrice     = 2.0
	buyingPrice      = 10.0
	resourceObjKey   = "resource/gold"
	resourceKind     = "gold"
	itemToSellObjKey = "resource/log"
	itemToSellKind   = "log"
	itemToBuyObjKey  = "resource/claim_stone"
	itemToBuyKind    = "claim_stone"
)

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

func testPlayerNeedToBeCloser(t *testing.T) {
	res, err := testFunction()
	assert.False(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, "Player need to be closer to NPC", err.Error())
}

func testPlayerDoesNotHaveRequiredResources(t *testing.T) {
	res, err := testFunction()
	assert.False(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, "Player does not have required resources", err.Error())
}

func testPlayerCanNotRemoveRequiredResources(t *testing.T) {
	res, err := testFunction()
	assert.False(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, "Player can not remove required resources", err.Error())
}

func testSuccess(t *testing.T) {
	res, err := testFunction()
	assert.True(t, res)
	assert.Nil(t, err)
}

func callBuyItem() (bool, error) {
	return npcObj.(entity.INpcObject).BuyItem(game_test.GetEngine(), charGameObj, itemToBuyObjKey, amount)
}

func callSellItem() (bool, error) {
	return npcObj.(entity.INpcObject).SellItem(game_test.GetEngine(), charGameObj, itemToSellObjKey, amount)
}
