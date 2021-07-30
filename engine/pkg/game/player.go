package game

type Player struct {
	Id int // equals to Character.Id
	CharacterGameObjectId int
	VisionAreaGameObjectId int
	Client *Client
	VisibleObjects []int
}
