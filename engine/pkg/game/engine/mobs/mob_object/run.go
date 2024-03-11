package mob_object

func (mob *MobObject) Run(newTickTime int64) {
	// Mob logic processing goes here:
	// Stop moving.
	if mob.State == MovingState && (newTickTime - mob.TickTime) >= MovingTime {
		mob.stop()
		mob.TickTime = newTickTime
	} else
	// Start random moving.
	if mob.State == IdleState && (newTickTime - mob.TickTime) >= IdleTime {
		mob.moveInRandomDirection()
		mob.TickTime = newTickTime
	} else
	// Stop following
	if mob.State == StopFollowingState {
		mob.TargetObjectId = ""
		mob.stop()
		mob.TickTime = newTickTime
	} else 
	// Start Following
	if mob.State == StartFollowState {
		mob.State = FollowingState
		mob.TickTime = newTickTime
		mob.setMoveTo(FollowingDirectionChangeTime)
	} else
	// Perform following
	if mob.State == FollowingState {
		if (newTickTime - mob.TickTime) >= FollowingTime {
			mob.Unfollow()
		} else { // Perform actual following
			if targetObj, ok := mob.Engine.GameObjects().Load(mob.TargetObjectId); ok {
				mob.performFollowing(targetObj, FollowingDirectionChangeTime)
			} else {
				mob.Unfollow()
			}
		}
	} else
	// Start Attacking
	if mob.State == StartAttackingState {
		mob.State = AttackingState
		mob.TickTime = newTickTime
		mob.setMoveTo(AttackingDirectionChangeTime)
		mob.Properties()["speed"] = mob.Properties()["speed"].(float64) * AttackSpeedUp
	} else
	// Renew Attacking
	if mob.State == RenewAttackingState {
		mob.State = AttackingState
		mob.TickTime = newTickTime
		mob.setMoveTo(AttackingDirectionChangeTime)
	} else
	// Stop attacking
	if mob.State == StopAttackingingState {
		mob.TargetObjectId = ""
		mob.Properties()["speed"] = mob.Properties()["speed"].(float64) / AttackSpeedUp
		mob.stop()
		mob.TickTime = newTickTime
	} else
	// Perform attacking
	if mob.State == AttackingState {
		if (newTickTime - mob.TickTime) >= AttackingTime {
			mob.StopAttacking()
		} else { // Perform actual following before hit
			if targetObj, ok := mob.Engine.GameObjects().Load(mob.TargetObjectId); ok {
				mob.performFollowing(targetObj, AttackingDirectionChangeTime)
				mob.MeleeHit(targetObj)
			} else {
				mob.StopAttacking()
			}
		}
	} 
}
