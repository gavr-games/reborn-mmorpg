package mob_object

func (mob *MobObject) StopAttacking() {
	mob.State = StopAttackingingState
}