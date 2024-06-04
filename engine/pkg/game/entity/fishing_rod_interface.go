package entity

type IFishingRodObject interface {
	CheckCatch(e IEngine, charGameObj IGameObject) (bool, error)
	Catch(e IEngine, charGameObj IGameObject) (bool, error)
}
