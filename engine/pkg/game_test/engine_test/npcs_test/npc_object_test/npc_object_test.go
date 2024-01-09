package npc_object_test

import (
	"os"
	"testing"

	"github.com/gavr-games/reborn-mmorpg/pkg/game"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/npcs/npc_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/stretchr/testify/assert"
)

var e *game.Engine = game.NewEngine()
var npcObj entity.IGameObject
var charGameObj entity.IGameObject
var testFunction func() (bool, error)

const (
	amount           = 1.0
	itemToSellObjKey = "resource/log"
	itemToBuyObjKey  = "resource/claim_stone"
)

func TestMain(m *testing.M) {
	e.Init()
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
	return npcObj.(*npc_object.NpcObject).BuyItem(e, charGameObj, itemToBuyObjKey, amount)
}

func callSellItem() (bool, error) {
	return npcObj.(*npc_object.NpcObject).SellItem(e, charGameObj, itemToSellObjKey, amount)
}
