package mob_object

import (
	"context"
	"pgregory.net/rand"
	"sync"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/leveling_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/moving_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/melee_weapon_object"

	"github.com/looplab/fsm"
)

const (
	IdleTime = 80000.0 // stays idle during this time
	MovingTime = 5000.0 // randomly moves during this time
	FollowingTime = 40000.0 // stops following after this time
	FollowingDistance = 0.2 // stops when in range of the target
	FollowingDirectionChangeTime = 500.0 // change direction only once per this time
	AttackSpeedUp = 1.5 // increases the mob speed during attack
	AttackingTime = 20000.0 // during this time the mob attacks until it calms down if not hitted back
	AttackingDirectionChangeTime = 500.0 // change direction only once per this time
	RecentlyKilledTime = 2000.0 // stop attacking if target was recently killed
	AgressiveCheckProbability = 0.05 // allows not to check agressive every game cycle
	ControlRange = 20.0 // the distance to control your mob
)

// TODO: refactor to thread safe
type MobObject struct {
	Engine         entity.IEngine
	TickTime       int64
	TargetObjectId string
	FSM            *fsm.FSM
	moving_object.MovingObject
	leveling_object.LevelingObject
	melee_weapon_object.MeleeWeaponObject
	entity.GameObject
	propsMutex     sync.RWMutex
}

func (mob *MobObject) GetTickTime() int64 {
	mob.propsMutex.RLock()
	defer mob.propsMutex.RUnlock()
	return mob.TickTime
}

func (mob *MobObject) SetTickTime(newTickTime int64) {
	mob.propsMutex.Lock()
	defer mob.propsMutex.Unlock()
	mob.TickTime = newTickTime
}

func (mob *MobObject) GetTargetObjectId() string {
	mob.propsMutex.RLock()
	defer mob.propsMutex.RUnlock()
	return mob.TargetObjectId
}

func (mob *MobObject) SetTargetObjectId(targetObjectId string) {
	mob.propsMutex.Lock()
	defer mob.propsMutex.Unlock()
	mob.TargetObjectId = targetObjectId
}

func NewMobObject(e entity.IEngine, gameObj entity.IGameObject) *MobObject {
	mob := &MobObject{
		e,
		e.CurrentTickTime() + rand.Int63n(IdleTime),
		"", // for following and attack
		nil,
		moving_object.MovingObject{},
		leveling_object.LevelingObject{},
		melee_weapon_object.MeleeWeaponObject{},
		*gameObj.(*entity.GameObject),
		*new(sync.RWMutex),
	}
	mob.InitMovingObject(mob)
	mob.InitLevelingObject(mob)
	mob.InitMeleeWeaponObject(mob)

	mob.SetupFSM()

	return mob
}

func (mob *MobObject) SetupFSM() {
	mob.FSM = fsm.NewFSM(
		"idle",
		fsm.Events{
			{Name: "move", Src: []string{"idle"}, Dst: "moving"},
			{Name: "follow", Src: []string{"idle", "moving", "following", "attacking"}, Dst: "following"},
			{Name: "attack", Src: []string{"idle", "moving", "following", "attacking"}, Dst: "attacking"},
			{Name: "stop", Src: []string{"idle", "moving", "following", "attacking"}, Dst: "idle"},
			// {Name: "die", Src: []string{"idle", "moving", "following", "attacking"}, Dst: "dead"},
			// {Name: "resurrect", Src: []string{"dead"}, Dst: "idle"},
		},
		fsm.Callbacks{
			"move": func(_ context.Context, e *fsm.Event) {
				mob.moveInRandomDirection()
				mob.SetTickTime(mob.Engine.CurrentTickTime())
			},
			"stop": func(_ context.Context, e *fsm.Event) {
				mob.Stop(mob.Engine)
				mob.SetTickTime(mob.Engine.CurrentTickTime())
			},
			"follow": func(ctx context.Context, e *fsm.Event) {
				mob.SetTargetObjectId(ctx.Value("targetObjId").(string))
				mob.SetTickTime(mob.Engine.CurrentTickTime())
				mob.setMoveTo(FollowingDirectionChangeTime)
			},
			"leave_following": func(_ context.Context, e *fsm.Event) {
				mob.SetTargetObjectId("")
			},
			"attack": func(ctx context.Context, e *fsm.Event) {
				mob.SetTargetObjectId(ctx.Value("targetObjId").(string))
				mob.SetTickTime(mob.Engine.CurrentTickTime())
				mob.setMoveTo(AttackingDirectionChangeTime)
			},
			"enter_attacking": func(ctx context.Context, e *fsm.Event) {
				mob.SetProperty("speed", mob.GetProperty("speed").(float64) * AttackSpeedUp)
			},
			"leave_attacking": func(ctx context.Context, e *fsm.Event) {
				mob.SetTargetObjectId("")
				mob.SetProperty("speed", mob.GetProperty("speed").(float64) / AttackSpeedUp)
			},
		},
	)
}
