package entity

type ILevelingObject interface {
	AddExperience(e IEngine, action string) (bool, error)
}
