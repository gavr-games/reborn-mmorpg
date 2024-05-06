package entity

type INpcObject interface {
	BuyItem(e IEngine, charGameObj IGameObject, itemKey string, amount float64) (bool, error)
	SellItem(e IEngine, charGameObj IGameObject, itemKey string, amount float64) (bool, error)
	GetDungeonsInfo(e IEngine, charGameObj IGameObject) (map[string]interface{}, error)
}
