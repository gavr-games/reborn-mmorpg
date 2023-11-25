package constants

const (
	FloorSize = 200.0
	FloorCount = 4
	TownSize = 9.0
	TickSize = 10 // Game tick size in ms
	InitialPlayerX = 100.0
	InitialPlayerY = 100.0
	PlayerVisionArea = 70.0
	PlayerSpeed = 2.0
	ClaimArea = 20.0
	ClaimRentDuration = 2419200000.0 // 4 weeks
	ClaimRentCost = 10.0 //gold
)

func GetPossibleDirections() []string {
	// Order is important here for finding direction by angle in mobs/mob.go
	return []string {"move_east", "move_north_east", "move_north", "move_north_west", "move_west", "move_south_west", "move_south", "move_south_east"}
}
