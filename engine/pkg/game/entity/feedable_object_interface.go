package entity

type IFeedableObject interface {
	Feed(e IEngine, itemId string, player *Player) (bool, error)
}
