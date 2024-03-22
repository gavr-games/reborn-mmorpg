package entity

type IClaimObeliskObject interface {
	Destroy(e IEngine, player *Player) bool
	ExtendRent(e IEngine) bool
	Init(e IEngine) bool
	Remove(e IEngine) bool
	FindKindInArea(e IEngine, kind string) (IGameObject, error)
}
