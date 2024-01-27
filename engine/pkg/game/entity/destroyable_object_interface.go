package entity

type IDestroyableObject interface {
	Destroy(e IEngine, player *Player) bool
}
