package entity

type Player struct {
	Id int // equals to Character.Id
	CharacterGameObjectId string
	VisionAreaGameObjectId string
	Client IClient
	VisibleObjects []string
}
