package mob_object

func (mob *MobObject) Die() {
	mob.drop()

	// Remove from world
	mob.Engine.RemoveGameObject(mob)
}