package dragon_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/feedable_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/evolvable_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mobs/mob_object"
)

const (
	RESURRECT_COST_PER_LEVEL = 25
)

type DragonObject struct {
	evolvable_object.EvolvableObject
	feedable_object.FeedableObject
	mob_object.MobObject
}

func NewDragonObject(e entity.IEngine, gameObj entity.IGameObject) *DragonObject {
	dragon := &DragonObject{
		evolvable_object.EvolvableObject{},
		feedable_object.FeedableObject{},
		*mob_object.NewMobObject(e, gameObj),
	}
	dragon.InitEvolvableObject(dragon)
	dragon.InitFeedableObject(dragon)
	dragon.SetupFSM() // we need to call this again, so FSM callbacks get updated pointer to dragon

	return dragon
}
