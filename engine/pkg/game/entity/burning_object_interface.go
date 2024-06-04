package entity

type IBurningObject interface {
	Burn(e IEngine, player *Player) (bool, error)
	Extinguish(e IEngine, player *Player) (bool, error)
	AddFuel(e IEngine, itemId string, player *Player) (bool, error)
}
