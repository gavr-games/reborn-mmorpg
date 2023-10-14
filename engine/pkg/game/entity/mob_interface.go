package entity

type MobCommand struct {
	// command to perform
	Command int

	// command params
	Params map[string]interface{}
}

type IMob interface {
	GetId() string
	Run(newTickTime int64)
	Follow(targetObjId string)
	Unfollow()
}
