package entity

type IMobObject interface {
	Run(newTickTime int64)
	Follow(targetObjId string)
	Unfollow()
	Die()
	Attack(targetObjId string)
	StopAttacking()
}
