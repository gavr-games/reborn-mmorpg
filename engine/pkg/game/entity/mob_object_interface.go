package entity

type IMobObject interface {
	Run(newTickTime int64)
	OrderToAttack(ownerObj IGameObject)
	OrderToFollow(targetObj IGameObject)
	OrderToStop(ownerObj IGameObject)
	Die()
	Attack(targetObjId string)
	StopEverything()
	CheckAgro()
}
