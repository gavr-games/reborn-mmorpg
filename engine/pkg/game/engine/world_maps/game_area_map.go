package world_maps

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

const (
	MapsFolder = "/src/github.com/gavr-games/reborn-mmorpg/assets/maps"
)

var surfaceColors = map[string]color.RGBA{
	"surface/dirt": color.RGBA{165, 42, 42, 255},
	"surface/grass": color.RGBA{0, 255, 0, 255},
	"surface/sand": color.RGBA{255, 255, 0, 255},
	"surface/stone": color.RGBA{128, 128, 128, 255},
	"surface/stone_road": color.RGBA{100, 100, 100, 255},
	"surface/town_floor": color.RGBA{128, 128, 128, 255},
	"surface/water": color.RGBA{0, 0, 255, 255},
}

type WorldCell struct {
	X, Y float64
	SurfaceKind string
}

// Generates JPEG image with GameArea map during world generation
type GameAreaMap struct {
	Img      *image.RGBA
	GameArea *entity.GameArea
	Cells    chan *WorldCell
	Finish   chan bool
}

func NewGameAreaMap(gameArea *entity.GameArea) *GameAreaMap {
	gaMap := &GameAreaMap{
		Img: image.NewRGBA(image.Rect(int(gameArea.X()), int(gameArea.Y()), int(gameArea.Width()), int(gameArea.Height()))),
		GameArea: gameArea,
		Cells: make(chan *WorldCell),
		Finish: make(chan bool),
	}
	draw.Draw(gaMap.Img, gaMap.Img.Bounds(), &image.Uniform{surfaceColors["surface/grass"]}, image.ZP, draw.Src)

	go gaMap.Run()

	return gaMap
}

func (gaMap *GameAreaMap) Run() {
	defer close(gaMap.Cells)
	defer close(gaMap.Finish)
	for {
		select {
		case cell := <-gaMap.Cells:
			// draw cell
			bounds := image.Rect(int(cell.X), int(gaMap.GameArea.Width()) - 1 - int(cell.Y), int(cell.X) + 1, int(gaMap.GameArea.Height()) - int(cell.Y))
			draw.Draw(gaMap.Img, bounds, &image.Uniform{surfaceColors[cell.SurfaceKind]}, image.Point{int(cell.X), int(gaMap.GameArea.Height()) - 1 - int(cell.Y)}, draw.Src)
		case finish := <-gaMap.Finish:
			if finish {
				filename := fmt.Sprintf("%s/area_%s_map.jpg", MapsFolder, gaMap.GameArea.Id())
				f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
				if err != nil {
						panic(err)
				}
				defer f.Close()
				if err = jpeg.Encode(f, gaMap.Img, &jpeg.Options{100}); err != nil {
					panic(err)
				}
				gaMap.Img = nil
				break
			}
		}
	}
}
