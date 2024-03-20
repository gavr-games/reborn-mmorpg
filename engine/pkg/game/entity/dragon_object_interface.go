package entity

type IDragonObject interface {
	TeleportToOwner(charGameObj IGameObject)
	Release(charGameObj IGameObject)
}
