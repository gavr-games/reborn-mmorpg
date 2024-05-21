package dungeons

import (
	"math/rand"
)

const (
	SizeMultiplier = 1.5
)

// Use BSP Dungeon generation https://www.roguebasin.com/index.php/Basic_BSP_Dungeon_generation
func BSPGenerate(width, height, minRoomSize, maxRoomSize float64, splitsCount int) *Container {
	rootContainer := &Container{
		0.0,
		0.0,
		width,
		height,
		nil, nil, nil,
	}
	splitContainer(rootContainer, minRoomSize, maxRoomSize, splitsCount)
	return rootContainer
}

func splitContainer(cont *Container, minRoomSize, maxRoomSize float64, splitsCount int) {
	if splitsCount == 0 {
		return
	}

	maxSize := maxRoomSize * SizeMultiplier
	if cont.Width >= maxSize && cont.Height < maxSize {
		splitVertically(cont, minRoomSize)
	} else if cont.Height >= maxSize && cont.Width < maxSize {
		splitHorizontally(cont, minRoomSize)
	} else if cont.Height >= maxSize && cont.Width >= maxSize {
		if rand.Float64() < 0.5 {
			splitVertically(cont, minRoomSize)
		} else {
			splitHorizontally(cont, minRoomSize)
		}
	}

	if cont.Child1 != nil {
		splitContainer(cont.Child1, minRoomSize, maxRoomSize, splitsCount - 1)
	}
	if cont.Child2 != nil {
		splitContainer(cont.Child2, minRoomSize, maxRoomSize, splitsCount - 1)
	}
}

func splitHorizontally(cont *Container, minRoomSize float64) {
	min := minRoomSize
  max := cont.Height - minRoomSize
	splitY := float64(rand.Intn(int(max - min))) + min
	cont.Child1 = &Container{
		cont.X,
		cont.Y,
		cont.Width,
		splitY,
		nil, nil, nil,
	}
	cont.Child2 = &Container{
		cont.X,
		cont.Y + splitY,
		cont.Width,
		cont.Height - splitY,
		nil, nil, nil,
	}
}

func splitVertically(cont *Container, minRoomSize float64) {
	min := minRoomSize
  max := cont.Width - minRoomSize
	splitX := float64(rand.Intn(int(max - min))) + min
	cont.Child1 = &Container{
		cont.X,
		cont.Y,
		splitX,
		cont.Height,
		nil, nil, nil,
	}
	cont.Child2 = &Container{
		cont.X + splitX,
		cont.Y,
		cont.Width - splitX,
		cont.Height,
		nil, nil, nil,
	}
}
