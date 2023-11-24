package entity

type ICactusObject interface {
	CheckCut(e IEngine, charGameObj IGameObject) bool
	Cut(e IEngine, charGameObj IGameObject) bool
}
