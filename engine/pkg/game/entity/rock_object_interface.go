package entity

type IRockObject interface {
	CheckChip(e IEngine, charGameObj IGameObject) bool
	Chip(e IEngine, charGameObj IGameObject) bool
}
