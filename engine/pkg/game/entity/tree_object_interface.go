package entity

type ITreeObject interface {
	CheckChop(e IEngine, charGameObj IGameObject) bool
	Chop(e IEngine, charGameObj IGameObject) bool
}
