package mob_object

func (mob *MobObject) Attack(targetObjId string) {
	if mob.State == AttackingState {
		mob.State = RenewAttackingState
	} else {
		mob.State = StartAttackingState
		mob.TargetObjectId = targetObjId
	}
}