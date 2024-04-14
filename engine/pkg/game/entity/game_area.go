package entity

import (
	"encoding/json"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/satori/go.uuid"
	"go.uber.org/atomic"
)

const (
	DefaultMaxObjects = 200
	DefaultMaxLevels = 10
)

// Struct to hold different areas/territories of th game: dungeoun, surface, underground...
type GameArea struct {
	id     *atomic.String
	key    *atomic.String
	x      *atomic.Float64
	y      *atomic.Float64
	width  *atomic.Float64
	height *atomic.Float64
	utils.Quadtree
}

func (ga *GameArea) Id() string {
	return ga.id.Load()
}

func (ga *GameArea) SetId(id string) {
	ga.id.Store(id)
}

func (ga *GameArea) Key() string {
	return ga.key.Load()
}

func (ga *GameArea) SetKey(key string) {
	ga.key.Store(key)
}

func (ga *GameArea) X() float64 {
	return ga.x.Load()
}

func (ga *GameArea) SetX(x float64) {
	ga.x.Store(x)
}

func (ga *GameArea) Y() float64 {
	return ga.y.Load()
}

func (ga *GameArea) SetY(y float64) {
	ga.x.Store(y)
}

func (ga *GameArea) Width() float64 {
	return ga.width.Load()
}

func (ga *GameArea) SetWidth(width float64) {
	ga.width.Store(width)
}

func (ga *GameArea) Height() float64 {
	return ga.height.Load()
}

func (ga *GameArea) SetHeight(height float64) {
	ga.height.Store(height)
}

func NewGameArea(key string, x, y, width, height float64) *GameArea {
	id := uuid.NewV4().String()
	ga := &GameArea{
		atomic.NewString(id),
		atomic.NewString(key),
		atomic.NewFloat64(x),
		atomic.NewFloat64(y),
		atomic.NewFloat64(width),
		atomic.NewFloat64(height),
		utils.Quadtree{},
	}
	ga.InitQuadtree()
	return ga
}

func (ga *GameArea) InitQuadtree() {
	ga.Quadtree = utils.Quadtree{
		Bounds: utils.Bounds{
			X:      ga.X(),
			Y:      ga.Y(),
			Width:  ga.Width(),
			Height: ga.Height(),
		},
		MaxObjects: DefaultMaxObjects,
		MaxLevels:  DefaultMaxLevels,
		Level:      0,
		Objects:    make([]utils.IBounds, 0),
		Nodes:      make([]utils.Quadtree, 0),
	}
}

func (ga *GameArea) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Id     string
		Key    string
		X      float64
		Y      float64
		Width  float64
		Height float64
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}
	ga.id = atomic.NewString(tmp.Id)
	ga.key = atomic.NewString(tmp.Key)
	ga.x = atomic.NewFloat64(tmp.X)
	ga.y = atomic.NewFloat64(tmp.Y)
	ga.width = atomic.NewFloat64(tmp.Width)
	ga.height = atomic.NewFloat64(tmp.Height)
	ga.InitQuadtree()
	return nil
}

func (ga *GameArea) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id     string
		Key    string
		X      float64
		Y      float64
		Width  float64
		Height float64
	}{
		Id:     ga.Id(),
		Key:    ga.Key(),
		X:      ga.X(),
		Y:      ga.Y(),
		Width:  ga.Width(),
		Height: ga.Height(),
	})
}
