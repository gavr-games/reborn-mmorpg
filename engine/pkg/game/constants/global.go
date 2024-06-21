package constants

const (
	SurfaceSize = 200.0 // size of initial GameArea
	TickSize = 10 // Game tick size in ms
	InitialPlayerArea = "town"
	InitialPlayerX = 25.0
	InitialPlayerY = 25.0
	PlayerVisionArea = 60.0
	PlayerSpeed = 3.0
	ClaimArea = 20.0
	ClaimRentDuration = 2419200000.0 // 4 weeks
	ClaimRentCost = 10.0 //gold
	MoveToDefaultDirectionChangeTime = 1000.0 // ms, used for MoveToCoords feature
)

func GetPossibleDirections() []string {
	// Order is important here for finding direction by angle in mobs/mob.go
	return []string {"move_east", "move_north_east", "move_north", "move_north_west", "move_west", "move_south_west", "move_south", "move_south_east"}
}
