package mob_object

func (mob *MobObject) Follow(targetObjId string) {
	mob.State = StartFollowState
	mob.TargetObjectId = targetObjId
}
