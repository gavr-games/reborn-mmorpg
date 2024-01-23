package entity

type IPlantObject interface {
	CheckCut(e IEngine, charGameObj IGameObject) bool
	Cut(e IEngine, charGameObj IGameObject) bool
	CheckHarvest(e IEngine, charGameObj IGameObject) bool
	Harvest(e IEngine, charGameObj IGameObject) bool
	Grow(e IEngine) bool
}
