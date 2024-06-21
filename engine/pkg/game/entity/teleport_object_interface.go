package entity

type ITeleportObject interface {
	TeleportTo(e IEngine, charGameObj IGameObject) (bool, error)
}
