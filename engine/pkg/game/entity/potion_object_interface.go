package entity

type IPotionObject interface {
	ApplyToPlayer(e IEngine, player *Player) bool
}
