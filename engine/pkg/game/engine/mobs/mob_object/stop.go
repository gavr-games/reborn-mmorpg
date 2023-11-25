package mob_object

func (mob *MobObject) stop() {
	mob.Stop(mob.Engine)
	mob.State = IdleState
}
