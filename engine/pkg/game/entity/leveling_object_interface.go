package entity

type ILevelingObject interface {
	AddExperience(e IEngine, action string) (bool, error)
	AddDungeonExperience(e IEngine, level float64) (bool, error)
}
