package entity

type IShovelObject interface {
	CheckDig(e IEngine, charGameObj IGameObject) bool
	Dig(e IEngine, charGameObj IGameObject) bool
}
