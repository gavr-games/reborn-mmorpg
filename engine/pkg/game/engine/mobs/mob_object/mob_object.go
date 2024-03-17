package mob_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/leveling_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/moving_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/melee_weapon_object"
)

const (
	IdleState = 0
	MovingState = 1
	StartFollowState = 2
	FollowingState = 3
	StopFollowingState = 4
	StartAttackingState = 5
	AttackingState = 6
	StopAttackingingState = 7
	RenewAttackingState = 8
)

const (
	IdleTime = 40000.0 // stays idle during this time
	MovingTime = 5000.0 // randomly moves during this time
	FollowingTime = 40000.0 // stops following after this time
	FollowingDistance = 0.2 // stops when in range of the target
	FollowingDirectionChangeTime = 500.0 // change direction only once per this time
	AttackSpeedUp = 1.5 // increases the mob speed during attack
	AttackingTime = 20000.0 // during this time the mob attacks until it calms down if not hitted back
	AttackingDirectionChangeTime = 500.0 // change direction only once per this time
)

type MobObject struct {
	Engine entity.IEngine
	TickTime int64
	State int
	TargetObjectId string
	moving_object.MovingObject
	leveling_object.LevelingObject
	melee_weapon_object.MeleeWeaponObject
	entity.GameObject
}

func NewMobObject(e entity.IEngine, gameObj entity.IGameObject) *MobObject {
	mob := &MobObject{
		e,
		e.CurrentTickTime(),
		IdleState,
		"", // for following and attack
		moving_object.MovingObject{},
		leveling_object.LevelingObject{},
		melee_weapon_object.MeleeWeaponObject{},
		*gameObj.(*entity.GameObject),
	}
	mob.InitMovingObject(mob)
	mob.InitLevelingObject(mob)
	mob.InitMeleeWeaponObject(mob)

	return mob
}
