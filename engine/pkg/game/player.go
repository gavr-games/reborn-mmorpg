package game

type Player struct {
	Id int // equals to Character.Id
	GameObjectId int
	Client *Client
	VisibleObjects []int
}
