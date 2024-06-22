package utils

import (
	"sync"
)

// Quadtree - The quadtree data structure
type Quadtree struct {
	Bounds     Bounds
	MaxObjects int // Maximum objects a node can hold before splitting into 4 subnodes
	MaxLevels  int // Total max levels inside root Quadtree
	Level      int // Depth level, required for subnodes
	Objects    []IBounds
	Nodes      []Quadtree
	Total      int
	mu         sync.RWMutex
}

// TotalNodes - Retrieve the total number of sub-Quadtrees in a Quadtree
func (qt *Quadtree) TotalNodes() int {

	total := 0

	if len(qt.Nodes) > 0 {
		for i := 0; i < len(qt.Nodes); i++ {
			total += 1
			total += qt.Nodes[i].TotalNodes()
		}
	}

	return total

}

// split - split the node into 4 subnodes
func (qt *Quadtree) split() {

	if len(qt.Nodes) == 4 {
		return
	}

	nextLevel := qt.Level + 1
	subWidth := qt.Bounds.Width / 2
	subHeight := qt.Bounds.Height / 2
	x := qt.Bounds.X
	y := qt.Bounds.Y

	//top right node (0)
	qt.Nodes = append(qt.Nodes, Quadtree{
		Bounds: Bounds{
			X:      x + subWidth,
			Y:      y,
			Width:  subWidth,
			Height: subHeight,
		},
		MaxObjects: qt.MaxObjects,
		MaxLevels:  qt.MaxLevels,
		Level:      nextLevel,
		Objects:    make([]IBounds, 0),
		Nodes:      make([]Quadtree, 0, 4),
	})

	//top left node (1)
	qt.Nodes = append(qt.Nodes, Quadtree{
		Bounds: Bounds{
			X:      x,
			Y:      y,
			Width:  subWidth,
			Height: subHeight,
		},
		MaxObjects: qt.MaxObjects,
		MaxLevels:  qt.MaxLevels,
		Level:      nextLevel,
		Objects:    make([]IBounds, 0),
		Nodes:      make([]Quadtree, 0, 4),
	})

	//bottom left node (2)
	qt.Nodes = append(qt.Nodes, Quadtree{
		Bounds: Bounds{
			X:      x,
			Y:      y + subHeight,
			Width:  subWidth,
			Height: subHeight,
		},
		MaxObjects: qt.MaxObjects,
		MaxLevels:  qt.MaxLevels,
		Level:      nextLevel,
		Objects:    make([]IBounds, 0),
		Nodes:      make([]Quadtree, 0, 4),
	})

	//bottom right node (3)
	qt.Nodes = append(qt.Nodes, Quadtree{
		Bounds: Bounds{
			X:      x + subWidth,
			Y:      y + subHeight,
			Width:  subWidth,
			Height: subHeight,
		},
		MaxObjects: qt.MaxObjects,
		MaxLevels:  qt.MaxLevels,
		Level:      nextLevel,
		Objects:    make([]IBounds, 0),
		Nodes:      make([]Quadtree, 0, 4),
	})

}

// getIndex - Determine which quadrant the object belongs to (0-3)
func (qt *Quadtree) getIndex(pRect IBounds) int {

	rect := pRect.HitBox()

	index := -1 // index of the subnode (0-3), or -1 if pRect cannot completely fit within a subnode and is part of the parent node

	verticalMidpoint := qt.Bounds.X + (qt.Bounds.Width / 2)
	horizontalMidpoint := qt.Bounds.Y + (qt.Bounds.Height / 2)

	//pRect can completely fit within the top quadrants
	topQuadrant := (rect.Y < horizontalMidpoint) && (rect.Y+rect.Height < horizontalMidpoint)

	//pRect can completely fit within the bottom quadrants
	bottomQuadrant := (rect.Y > horizontalMidpoint)

	//pRect can completely fit within the left quadrants
	if (rect.X < verticalMidpoint) && (rect.X+rect.Width < verticalMidpoint) {

		if topQuadrant {
			index = 1
		} else if bottomQuadrant {
			index = 2
		}

	} else if rect.X > verticalMidpoint {
		//pRect can completely fit within the right quadrants

		if topQuadrant {
			index = 0
		} else if bottomQuadrant {
			index = 3
		}

	}

	return index

}

// Insert - Insert the object into the node. If the node exceeds the capacity,
// it will split and add all objects to their corresponding subnodes.
func (qt *Quadtree) Insert(pRect IBounds) {
	qt.mu.Lock()
	defer qt.mu.Unlock()

	qt.Total++

	i := 0
	var index int

	// If we have subnodes within the Quadtree
	if len(qt.Nodes) > 0 == true {
		index = qt.getIndex(pRect)
		if index != -1 {
			qt.Nodes[index].Insert(pRect)
			return
		}
	}

	// If we don't subnodes within the Quadtree
	qt.Objects = append(qt.Objects, pRect)

	// If total objects is greater than max objects and level is less than max levels
	if (len(qt.Objects) > qt.MaxObjects) && (qt.Level < qt.MaxLevels) {
		// split if we don't already have subnodes
		if len(qt.Nodes) > 0 == false {
			qt.split()
		}
		// Add all objects to there corresponding subNodes
		for i < len(qt.Objects) {
			index = qt.getIndex(qt.Objects[i])
			if index != -1 {
				splice := qt.Objects[i]                                  // Get the object out of the slice
				qt.Objects = append(qt.Objects[:i], qt.Objects[i+1:]...) // Remove the object from the slice
				qt.Nodes[index].Insert(splice)
			} else {
				i++
			}
		}
	}
}

// Find object in quadtree via filter and remove it
func (qt *Quadtree) FilteredRemove(pRect IBounds, filter func(IBounds) bool) {
	qt.mu.Lock()
	defer qt.mu.Unlock()

	index := qt.getIndex(pRect)

	//if we have subnodes ...
	if len(qt.Nodes) > 0 {
		//if pRect fits into a subnode ..
		if index != -1 {
			qt.Nodes[index].FilteredRemove(pRect, filter)
		} else {
			// Find and remove item in current tree
			for i := 0; i < len(qt.Objects); i++ {
				result := filter(qt.Objects[i]) // check filtering condition
				if result {
					qt.Objects[i] = qt.Objects[len(qt.Objects) - 1] // Copy last element to index i.
					qt.Objects[len(qt.Objects) - 1] = nil // Erase last element (write zero value).
					qt.Objects = qt.Objects[:len(qt.Objects) - 1] // Truncate slice.
					break
				}
			}
		}
	} else {
		// Find and remove item in current tree
		for i := 0; i < len(qt.Objects); i++ {
			result := filter(qt.Objects[i]) // check filtering condition
			if result {
				qt.Objects[i] = qt.Objects[len(qt.Objects) - 1] // Copy last element to index i.
				qt.Objects[len(qt.Objects) - 1] = nil // Erase last element (write zero value).
				qt.Objects = qt.Objects[:len(qt.Objects) - 1] // Truncate slice.
				break
			}
		}
	}
}

// Find object in quadtree via filter and move it if possible
// returns false if pRect was removed and needs to be reinserted
func (qt *Quadtree) FilteredMove(pRect IBounds, newX, newY float64, filter func(IBounds) bool) bool {
	qt.mu.Lock()
	defer qt.mu.Unlock()

	index := qt.getIndex(pRect)

	//if we have subnodes ...
	if len(qt.Nodes) > 0 {
		//if pRect fits into a subnode ..
		if index != -1 {
			return qt.Nodes[index].FilteredMove(pRect, newX, newY, filter)
		} else {
			// Find and remove item in current tree
			for i := 0; i < len(qt.Objects); i++ {
				result := filter(qt.Objects[i]) // check filtering condition
				if result {
					maxX := qt.Bounds.X + qt.Bounds.Width
					maxY := qt.Bounds.Y + qt.Bounds.Height
					rect := pRect.HitBox()
					nexIndex := qt.getIndex(Bounds{ // Maybe moved to another subnode inside this node
						X:      newX,
						Y:      newY,
						Width:  rect.Width,
						Height: rect.Height,
					})
					// pRect is out of this quad tree node
					if index != nexIndex || newX + rect.Width > maxX || newX < qt.Bounds.X || newY + rect.Height > maxY || newY < qt.Bounds.Y {
						qt.Objects[i] = qt.Objects[len(qt.Objects) - 1] // Copy last element to index i.
						qt.Objects[len(qt.Objects) - 1] = nil // Erase last element (write zero value).
						qt.Objects = qt.Objects[:len(qt.Objects) - 1] // Truncate slice.
						return false
					}
					break
				}
			}
		}
	} else {
		// Find and remove item in current tree
		for i := 0; i < len(qt.Objects); i++ {
			result := filter(qt.Objects[i]) // check filtering condition
			if result {
				maxX := qt.Bounds.X + qt.Bounds.Width
				maxY := qt.Bounds.Y + qt.Bounds.Height
				rect := pRect.HitBox()
				// pRect is out of this quad tree node
				if newX + rect.Width > maxX || newX < qt.Bounds.X || newY + rect.Height > maxY || newY < qt.Bounds.Y {
					qt.Objects[i] = qt.Objects[len(qt.Objects) - 1] // Copy last element to index i.
					qt.Objects[len(qt.Objects) - 1] = nil // Erase last element (write zero value).
					qt.Objects = qt.Objects[:len(qt.Objects) - 1] // Truncate slice.
					return false
				}
				break
			}
		}
	}

	return true
}

// Retrieve - Return all objects that could collide with the given object
func (qt *Quadtree) Retrieve(pRect IBounds) []IBounds {
	index := qt.getIndex(pRect)
	// Array with all detected objects
	returnObjects := qt.Objects

	//if we have subnodes ...
	if len(qt.Nodes) > 0 {
		//if pRect fits into a subnode ..
		if index != -1 {
			returnObjects = append(returnObjects, qt.Nodes[index].Retrieve(pRect)...)
		} else {
			//if pRect does not fit into a subnode, check it against all subnodes
			for i := 0; i < len(qt.Nodes); i++ {
				if qt.Nodes[i].Bounds.Intersects(pRect.HitBox()) {
					returnObjects = append(returnObjects, qt.Nodes[i].Retrieve(pRect)...)
				}
			}
		}
	}

	return returnObjects
}

// RetrievePoints - Return all points that collide
func (qt *Quadtree) RetrievePoints(find IBounds) []IBounds {
	var foundPoints []IBounds
	potentials := qt.Retrieve(find)

	rect := find.HitBox()

	for o := 0; o < len(potentials); o++ {

		// X and Ys are the same and it has no Width and Height (Point)
		xyMatch := potentials[o].HitBox().X == float64(rect.X) && potentials[o].HitBox().Y == float64(rect.Y)
		if xyMatch && potentials[o].IsPoint() {
			foundPoints = append(foundPoints, find)
		}
	}

	return foundPoints
}

// RetrieveIntersections - Bring back all the bounds in a Quadtree that intersect with a provided bounds
func (qt *Quadtree) RetrieveIntersections(find IBounds) []IBounds {
	qt.mu.RLock()
	defer qt.mu.RUnlock()

	var foundIntersections []IBounds

	potentials := qt.Retrieve(find)
	for o := 0; o < len(potentials); o++ {
		if potentials[o].Intersects(find.HitBox()) {
			foundIntersections = append(foundIntersections, potentials[o])
		}
	}

	return foundIntersections
}

//Clear - Clear the Quadtree
func (qt *Quadtree) Clear() {
	qt.Objects = []IBounds{}

	if len(qt.Nodes)-1 > 0 {
		for i := 0; i < len(qt.Nodes); i++ {
			qt.Nodes[i].Clear()
		}
	}

	qt.Nodes = []Quadtree{}
	qt.Total = 0
}
