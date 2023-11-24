package entity

type IHatcheryObject interface {
	CheckHatch(e IEngine, charGameObj IGameObject) bool
	Hatch(e IEngine, mobPath string) bool
}
