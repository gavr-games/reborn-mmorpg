package entity

type IDragonObject interface {
	Die()
	TeleportToOwner(charGameObj IGameObject) (bool, error)
	Release(charGameObj IGameObject)
	Resurrect(charGameObj IGameObject) (bool, error)
}
