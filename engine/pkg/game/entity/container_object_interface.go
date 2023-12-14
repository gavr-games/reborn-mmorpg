package entity

type IContainerObject interface {
	CheckAccess(e IEngine, player *Player) bool
	GetItemKind(e IEngine, itemKind string) IGameObject
	GetItems(e IEngine) map[string]interface{}
	HasItemsKinds(e IEngine, items map[string]interface{}) bool
	Put(e IEngine, player *Player, itemId string, position int) bool
	RemoveItemsKinds(e IEngine, player *Player, items map[string]interface{}) bool
	AddResource(e IEngine, player *Player, resourceObj IGameObject, amount float64) bool
	Remove(e IEngine, player *Player, itemId string) bool
}
