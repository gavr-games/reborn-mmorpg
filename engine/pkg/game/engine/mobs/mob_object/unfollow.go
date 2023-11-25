package mob_object

func (mob *MobObject) Unfollow() {
	mob.State = StopFollowingState
}
