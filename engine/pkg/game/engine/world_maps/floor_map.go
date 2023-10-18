package world_maps

import (
	"os"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
)

const (
	MapsFolder = "/src/github.com/gavr-games/reborn-mmorpg/assets/maps"
)

var surfaceColors = map[string]color.RGBA{
	"surface/grass": color.RGBA{0, 255, 0, 255},
	"surface/sand": color.RGBA{255, 255, 0, 255},
	"surface/water": color.RGBA{0, 0, 255, 255},
}

type WorldCell struct {
	X, Y float64
	SurfaceKind string
}

// Generates JPEG image with floor map during world generation
type FloorMap struct {
	Img *image.RGBA
	FloorNumber int
	FloorSize int
	Cells chan *WorldCell
	Finish chan bool
}

func NewFloorMap(floorNumber, floorSize int) *FloorMap {
	floorMap := &FloorMap{
		Img: image.NewRGBA(image.Rect(0, 0, floorSize, floorSize)),
		FloorNumber: floorNumber,
		FloorSize: floorSize,
		Cells: make(chan *WorldCell),
		Finish: make(chan bool),
	}
	draw.Draw(floorMap.Img, floorMap.Img.Bounds(), &image.Uniform{surfaceColors["surface/grass"]}, image.ZP, draw.Src)

	go floorMap.Run()

	return floorMap
}

func (floorMap *FloorMap) Run() {
	defer close(floorMap.Cells)
	defer close(floorMap.Finish)
	for {
		select {
		case cell := <-floorMap.Cells:
			// draw cell
			bounds := image.Rect(int(cell.X), floorMap.FloorSize - 1 - int(cell.Y), int(cell.X) + 1, floorMap.FloorSize - int(cell.Y))
			draw.Draw(floorMap.Img, bounds, &image.Uniform{surfaceColors[cell.SurfaceKind]}, image.Point{int(cell.X), floorMap.FloorSize - 1 - int(cell.Y)}, draw.Src)
		case finish := <-floorMap.Finish:
			if finish {
				filename := fmt.Sprintf("%s/floor_%d_map.jpg", MapsFolder, floorMap.FloorNumber)
				f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
				if err != nil {
						panic(err)
				}
				defer f.Close()
				if err = jpeg.Encode(f, floorMap.Img, &jpeg.Options{100}); err != nil {
					panic(err)
				}
				floorMap.Img = nil
				break
			}
		}
	}
}
