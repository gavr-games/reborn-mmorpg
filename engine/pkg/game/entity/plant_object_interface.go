package entity

type IPlantObject interface {
	CheckCut(e IEngine, charGameObj IGameObject) bool
	Cut(e IEngine, charGameObj IGameObject) bool
}
