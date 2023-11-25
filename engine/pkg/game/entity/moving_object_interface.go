package entity

type IMovingObject interface {
	CanMove(e IEngine, dx float64, dy float64) (float64, float64)
	SetXYSpeeds(e IEngine, direction string)
	Stop(e IEngine)
}
