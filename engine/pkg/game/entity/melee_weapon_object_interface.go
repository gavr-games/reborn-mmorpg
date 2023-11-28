package entity

type IMeleeWeaponObject interface {
	CanHit(attacker IGameObject, target IGameObject) bool
}
