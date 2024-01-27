package entity

type IPickableObject interface {
	Drop(e IEngine, player *Player) bool
	Pickup(e IEngine, player *Player) bool
	PutToContainer(e IEngine, containerId string, pos int, player *Player) bool
}
