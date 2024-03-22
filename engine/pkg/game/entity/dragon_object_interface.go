package entity

type IDragonObject interface {
	Die()
	TeleportToOwner(charGameObj IGameObject)
	Release(charGameObj IGameObject)
	Resurrect(charGameObj IGameObject) (bool, error)
}
