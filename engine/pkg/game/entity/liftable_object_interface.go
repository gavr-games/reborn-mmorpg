package entity

type ILiftableObject interface {
	Lift(e IEngine, charGameObj IGameObject) bool
	PutLifted(e IEngine, charGameObj IGameObject, x float64, y float64, rotation float64) bool
}
