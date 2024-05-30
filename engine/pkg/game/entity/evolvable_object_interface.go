package entity

type IEvolvableObject interface {
	Evolve(e IEngine, player *Player) (bool, error)
}
