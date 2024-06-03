package entity

type IDoorObject interface {
	Open(e IEngine, player *Player) (bool, error)
	Close(e IEngine, player *Player) (bool, error)
}
