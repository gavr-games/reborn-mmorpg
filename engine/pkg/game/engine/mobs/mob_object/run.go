package mob_object

import (
	"context"
)

// Mob logic processing goes here
func (mob *MobObject) Run(newTickTime int64) {
	// Stop moving.
	if mob.FSM.Current() == "moving" && (newTickTime - mob.GetTickTime()) >= MovingTime {
		mob.FSM.Event(context.Background(), "stop")
	} else
	// Start random moving.
	if mob.FSM.Current() == "idle" && (newTickTime - mob.GetTickTime()) >= IdleTime {
		mob.FSM.Event(context.Background(), "move")
	} else
	// Perform following
	if mob.FSM.Current() == "following" {
		if (newTickTime - mob.GetTickTime()) >= FollowingTime {
			mob.FSM.Event(context.Background(), "stop")
		} else { // Perform actual following
			if targetObj, ok := mob.Engine.GameObjects().Load(mob.GetTargetObjectId()); ok {
				mob.performFollowing(targetObj, FollowingDirectionChangeTime)
			} else {
				mob.FSM.Event(context.Background(), "stop")
			}
		}
	} else
	// Perform attacking
	if mob.FSM.Current() == "attacking" {
		if (newTickTime - mob.GetTickTime()) >= AttackingTime {
			mob.FSM.Event(context.Background(), "stop")
		} else { // Perform actual following before hit
			if targetObj, ok := mob.Engine.GameObjects().Load(mob.GetTargetObjectId()); ok {
				mob.performFollowing(targetObj, AttackingDirectionChangeTime)
				mob.MeleeHit(targetObj)
			} else {
				mob.FSM.Event(context.Background(), "stop")
			}
		}
	} 
}
