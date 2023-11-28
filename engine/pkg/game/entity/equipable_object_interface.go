package entity

type IEquipableObject interface {
	Equip(e IEngine, player *Player) bool
	Unequip(e IEngine, player *Player) bool
}
