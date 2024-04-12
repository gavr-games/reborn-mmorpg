package mob_object_test

import (
	"os"
	"testing"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mobs/mob_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test"
	"github.com/stretchr/testify/assert"
)

const (
	targetKey = "player/player"
	mobKey    = "mob/zombie"
	noAgroDistance = 100.0
)

var targetObj entity.IGameObject
var mob *mob_object.MobObject

func TestMain(m *testing.M) {
	game_test.Setup()
	os.Exit(m.Run())
}

func testNoCheckAgroTarget(t *testing.T) {
	mob.CheckAgro()
	targetObjId := mob.GetTargetObjectId()
	assert.Equal(t, "", targetObjId)
}

func testCheckAgroTarget(t *testing.T) {
	mob.CheckAgro()
	targetObjId := mob.GetTargetObjectId()
	assert.Equal(t, targetObj.Id(), targetObjId)
}
