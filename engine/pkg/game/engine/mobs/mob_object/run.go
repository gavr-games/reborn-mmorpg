package mob_object

import (
	"context"
	"math/rand"
)

// Mob logic processing goes here
func (mob *MobObject) Run(newTickTime int64) {
	// Check agressive
	if agressive := mob.GetProperty("agressive"); agressive != nil && agressive.(bool) && mob.FSM.Current() == "idle" {
		probability := rand.Float64()
		if probability <= AgressiveCheckProbability { // this is needed not to calculate agro on each game tick and improve performance
			go mob.CheckAgro()
		}
	}
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
				// Check someone recently killed the target
				if lastDeath := targetObj.GetProperty("last_death"); lastDeath != nil && (newTickTime - int64(lastDeath.(float64))) <= RecentlyKilledTime {
					mob.FSM.Event(context.Background(), "stop")
				} else {
					mob.performFollowing(targetObj, AttackingDirectionChangeTime)
					mob.MeleeHit(targetObj)
				}
			} else {
				mob.FSM.Event(context.Background(), "stop")
			}
		}
	} 
}
