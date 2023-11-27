package entity

type IBuildingObject interface {
	Destroy(e IEngine, player *Player) bool
}
