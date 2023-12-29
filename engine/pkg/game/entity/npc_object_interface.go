package entity

type INpcObject interface {
	BuyItem(e IEngine, charGameObj IGameObject, itemKey string, amount float64) bool
	SellItem(e IEngine, charGameObj IGameObject, itemKey string, amount float64) bool
}
