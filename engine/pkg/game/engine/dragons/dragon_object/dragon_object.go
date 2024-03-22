package dragon_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mobs/mob_object"
)

const (
	RESURRECT_COST_PER_LEVEL = 25
)

type DragonObject struct {
	mob_object.MobObject
}

func NewDragonObject(e entity.IEngine, gameObj entity.IGameObject) *DragonObject {
	dragon := &DragonObject{
		*mob_object.NewMobObject(e, gameObj),
	}

	return dragon
}
