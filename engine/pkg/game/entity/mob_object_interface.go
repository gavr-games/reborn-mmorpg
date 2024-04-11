package entity

type IMobObject interface {
	Run(newTickTime int64)
	Follow(targetObj IGameObject)
	Unfollow(targetObj IGameObject)
	Die()
	Attack(targetObjId string)
	StopEverything()
}
