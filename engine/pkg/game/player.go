package game

type Player struct {
	Id int // equals to Character.Id
	CharacterGameObjectId string
	VisionAreaGameObjectId string
	Client *Client
	VisibleObjects []string
}
