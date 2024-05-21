package dungeons

import (
	"math"
	"gonum.org/v1/gonum/stat/combin"
)

type Container struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
	Child1 *Container // sub container after division
	Child2 *Container	// sub sontainer after division
	Room   *Container // room of the dungeon generated inside container, only last containers in tree have rooms
}

func (c *Container) GetCenter() (float64, float64) {
	return c.X + c.Width / 2, c.Y + c.Height / 2
}

func (c *Container) GetAllRooms(rooms []*Container) []*Container  {
	if c.Room != nil {
		rooms = append(rooms, c.Room)
	}
	child1Rooms := []*Container{}
	if c.Child1 != nil {
		child1Rooms = c.Child1.GetAllRooms(rooms)
	}
	child2Rooms := []*Container{}
	if c.Child2 != nil {
		child2Rooms = c.Child2.GetAllRooms(rooms)
	}
	childRooms := append(child1Rooms, child2Rooms...)
	return append(rooms, childRooms...)
}

func (c *Container) GetAllCorridors(corridors []*Container, corridorSize float64) []*Container  {
	if c.Child1 != nil && c.Child2 != nil {
		child1X, child1Y := c.Child1.GetCenter()
		child2X, child2Y := c.Child2.GetCenter()
		child1X = math.Floor(child1X) - float64(int(math.Floor(child1X)) % 2)
		child2X = math.Floor(child2X) - float64(int(math.Floor(child2X)) % 2)
		child1Y = math.Floor(child1Y) - float64(int(math.Floor(child1Y)) % 2)
		child2Y = math.Floor(child2Y) - float64(int(math.Floor(child2Y)) % 2)
		if child2X > child1X { //horizontal corridor
			corridor := &Container{child1X, child1Y, child2X - child1X, corridorSize, nil, nil, nil}
			corridors = append(corridors, corridor)
		} else { //vertical corridor
			corridor := &Container{child1X, child1Y, corridorSize, child2Y - child1Y, nil, nil, nil}
			corridors = append(corridors, corridor)
		}
		corridors = append(corridors, c.Child1.GetAllCorridors(corridors, corridorSize)...)
		corridors = append(corridors, c.Child2.GetAllCorridors(corridors, corridorSize)...)
	}

	return corridors
}

// Return 3 rooms, which centers form maximum triangle area from combination of rooms https://pkg.go.dev/gonum.org/v1/gonum/stat/combin#Combinations
func (c *Container) GetThreeMostDistantRooms() []*Container {
	rooms := []*Container{}
	allRooms := []*Container{}
	allRooms = c.GetAllRooms(allRooms)
	maxTriangleArea := 0.0
	combs := combin.Combinations(len(allRooms), 3)
	for _, v := range combs {
		// Calculate traingle area
		r1X, r1Y := allRooms[v[0]].GetCenter()
		r2X, r2Y := allRooms[v[1]].GetCenter()
		r3X, r3Y := allRooms[v[2]].GetCenter()
		area := math.Abs(r1X * (r2Y - r3Y) + r2X * (r3Y - r1Y) + r3X * (r1Y - r2Y)) / 2
		if area > maxTriangleArea {
			maxTriangleArea = area
			rooms = []*Container{allRooms[v[0]], allRooms[v[1]], allRooms[v[2]]}
		}
	}

	return rooms
}
