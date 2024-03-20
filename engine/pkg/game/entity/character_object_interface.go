package entity

type ICharacterObject interface {
	HasTypeEquipped(e IEngine, itemType string) (IGameObject, bool)
	MeleeHit(e IEngine) bool
	Move(e IEngine, newX float64, newY float64, floor int)
	Reborn(e IEngine)
	SelectTarget(e IEngine, targetId string) bool
	DeselectTarget(e IEngine) bool
	GetVisionAreaX() float64
	GetVisionAreaY() float64
	GetDragonsInfo(e IEngine) map[string]interface{}
}
