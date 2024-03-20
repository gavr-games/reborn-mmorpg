package mob_object

func (mob *MobObject) StopEverything() {
	switch mob.State {
	case MovingState:
		mob.stop()
	case StartFollowState, FollowingState:
		mob.Unfollow(nil)
	case StartAttackingState, RenewAttackingState, AttackingState:
		mob.StopAttacking()
	}
}
