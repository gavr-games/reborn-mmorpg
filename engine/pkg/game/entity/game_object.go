package entity

import (
	"encoding/json"
	"math"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/constants"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

const (
	MaxDistance = 0.2
)

type IGameObject interface {
	X() float64
	SetX(x float64)
	Y() float64
	SetY(y float64)
	Width() float64
	SetWidth(width float64)
	Height() float64
	SetHeight(height float64)
	Id() string
	SetId(id string)
	Kind() string
	Type() string
	SetType(t string)
	Floor() int
	SetFloor(floor int)
	Rotation() float64
	SetRotation(rotation float64)
	CurrentAction() *DelayedAction
	SetCurrentAction(currentAction *DelayedAction)
	MoveToCoords() *MoveToCoords
	SetMoveToCoords(moveToCoords *MoveToCoords)
	SetMoveToCoordsByObject(moveToObj IGameObject)
	SetMoveToCoordsByXY(x float64, y float64)
	Properties() map[string]interface{}
	SetProperties(properties map[string]interface{})
	Effects() map[string]interface{}
	SetEffects(effects map[string]interface{})
	HitBox() utils.Bounds
	IsPoint() bool
	Intersects(b utils.Bounds) bool
	Clone() *GameObject
	GetDistance(b IGameObject) float64
	GetDistanceToXY(x float64, y float64) float64
	IsCloseTo(b IGameObject) bool
	Rotate(rotation float64)
	SetRotationByDirection(direction string)
	GetRotationByDirection(direction string) float64
	GetDirectionToXY(x float64, y float64) string
	TurnToXY(x float64, y float64) bool
}

type GameObject struct {
	// params for quadtree
	x      float64
	y      float64
	width  float64
	height float64

	// game params
	id            string
	objType       string
	floor         int // -1 for does not belong to any floor
	currentAction *DelayedAction
	rotation      float64 // from 0 to math.Pi * 2
	properties    map[string]interface{}
	effects       map[string]interface{}
	moveToCoords  *MoveToCoords //used for engine to automatically move object to this coord
}

func (obj *GameObject) X() float64 {
	return obj.x
}

func (obj *GameObject) SetX(x float64) {
	obj.x = x
	obj.properties["x"] = x
}

func (obj *GameObject) Y() float64 {
	return obj.y
}

func (obj *GameObject) SetY(y float64) {
	obj.y = y
	obj.properties["y"] = y
}

func (obj *GameObject) Width() float64 {
	return obj.width
}

func (obj *GameObject) SetWidth(width float64) {
	obj.width = width
	obj.properties["width"] = width
}

func (obj *GameObject) Height() float64 {
	return obj.height
}

func (obj *GameObject) SetHeight(height float64) {
	obj.height = height
	obj.properties["height"] = height
}

func (obj *GameObject) Id() string {
	return obj.id
}

func (obj *GameObject) SetId(id string) {
	obj.id = id
	obj.properties["id"] = id
}

func (obj *GameObject) Kind() string {
	return obj.properties["kind"].(string)
}

func (obj *GameObject) Type() string {
	return obj.objType
}

func (obj *GameObject) SetType(t string) {
	obj.objType = t
	obj.properties["type"] = t
}

func (obj *GameObject) Floor() int {
	return obj.floor
}

func (obj *GameObject) SetFloor(floor int) {
	obj.floor = floor
}

func (obj *GameObject) Rotation() float64 {
	return obj.rotation
}

func (obj *GameObject) SetRotation(rotation float64) {
	obj.rotation = rotation
}

func (obj *GameObject) CurrentAction() *DelayedAction {
	return obj.currentAction
}

func (obj *GameObject) SetCurrentAction(currentAction *DelayedAction) {
	obj.currentAction = currentAction
}

func (obj *GameObject) Properties() map[string]interface{} {
	return obj.properties
}

func (obj *GameObject) SetProperties(properties map[string]interface{}) {
	obj.properties = properties
}

func (obj *GameObject) Effects() map[string]interface{} {
	return obj.effects
}

func (obj *GameObject) SetEffects(effects map[string]interface{}) {
	obj.effects = effects
}

func (obj *GameObject) MoveToCoords() *MoveToCoords {
	return obj.moveToCoords
}

func (obj *GameObject) SetMoveToCoords(moveToCoords *MoveToCoords) {
	obj.moveToCoords = moveToCoords
}

func (obj *GameObject) SetMoveToCoordsByObject(moveToObj IGameObject) {
	obj.moveToCoords = &MoveToCoords{
		Mode: MoveCloseToBounds,
		Bounds: utils.Bounds{
			X:      moveToObj.X(),
			Y:      moveToObj.Y(),
			Width:  moveToObj.Width(),
			Height: moveToObj.Height(),
		},
		DirectionChangeTime:      constants.MoveToDefaultDirectionChangeTime,
		TimeUntilDirectionChange: 0,
	}
}

func (obj *GameObject) SetMoveToCoordsByXY(x float64, y float64) {
	obj.moveToCoords = &MoveToCoords{
		Mode: MoveToExactCoords,
		Bounds: utils.Bounds{
			X:      x,
			Y:      y,
			Width:  0.0,
			Height: 0.0,
		},
		DirectionChangeTime:      constants.MoveToDefaultDirectionChangeTime,
		TimeUntilDirectionChange: 0,
	}
}

func (obj *GameObject) UnmarshalJSON(b []byte) error {
	var tmp struct {
		X             float64
		Y             float64
		Width         float64
		Height        float64
		Id            string
		Type          string
		Floor         int
		CurrentAction *DelayedAction
		Rotation      float64
		Properties    map[string]interface{}
		Effects       map[string]interface{}
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}
	obj.x = tmp.X
	obj.y = tmp.Y
	obj.width = tmp.Width
	obj.height = tmp.Height
	obj.id = tmp.Id
	obj.objType = tmp.Type
	obj.floor = tmp.Floor
	obj.currentAction = tmp.CurrentAction
	obj.rotation = tmp.Rotation
	obj.properties = tmp.Properties
	obj.effects = tmp.Effects
	return nil
}

func (obj *GameObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		X             float64
		Y             float64
		Width         float64
		Height        float64
		Id            string
		Type          string
		Floor         int
		CurrentAction *DelayedAction
		Rotation      float64
		Properties    map[string]interface{}
		Effects       map[string]interface{}
	}{
		X:             obj.x,
		Y:             obj.y,
		Width:         obj.width,
		Height:        obj.height,
		Id:            obj.id,
		Type:          obj.objType,
		Floor:         obj.floor,
		CurrentAction: obj.currentAction,
		Rotation:      obj.rotation,
		Properties:    obj.properties,
		Effects:       obj.effects,
	})
}

func (obj *GameObject) HitBox() utils.Bounds {
	return utils.Bounds{
		X:      obj.X(),
		Y:      obj.Y(),
		Width:  obj.Width(),
		Height: obj.Height(),
	}
}

// IsPoint - Checks if a bounds object is a point or not (has no width or height)
func (obj *GameObject) IsPoint() bool {
	if obj.Width() == 0 && obj.Height() == 0 {
		return true
	}
	return false
}

// Intersects - Checks if a Bounds object intersects with another Bounds
func (a GameObject) Intersects(b utils.Bounds) bool {
	aMaxX := a.X() + a.Width()
	aMaxY := a.Y() + a.Height()
	bMaxX := b.X + b.Width
	bMaxY := b.Y + b.Height

	// a is left of b
	if aMaxX <= b.X {
		return false
	}

	// a is right of b
	if a.X() >= bMaxX {
		return false
	}

	// a is above b
	if aMaxY <= b.Y {
		return false
	}

	// a is below b
	if a.Y() >= bMaxY {
		return false
	}

	// The two overlap
	return true
}

func (obj *GameObject) Clone() *GameObject {
	clone := &GameObject{
		x:          obj.X(),
		y:          obj.Y(),
		width:      obj.Width(),
		height:     obj.Height(),
		id:         obj.Id(),
		objType:    obj.Type(),
		floor:      obj.Floor(),
		rotation:   obj.Rotation(),
		properties: make(map[string]interface{}),
		effects:    make(map[string]interface{}),
	}
	clone.SetProperties(utils.CopyMap(obj.Properties()))
	clone.SetEffects(utils.CopyMap(obj.Effects()))
	return clone
}

// Get approximate distance between objects. Assuming all of them are rectangles
func (a GameObject) GetDistance(b IGameObject) float64 {
	aXCenter := a.X() + a.Width()/2
	aYCenter := a.Y() + a.Height()/2

	bXCenter := b.X() + b.Width()/2
	bYCenter := b.Y() + b.Height()/2

	xDistance := math.Abs(aXCenter-bXCenter) - (a.Width()/2 + b.Width()/2)
	if xDistance < 0 {
		xDistance = 0.0
	}

	yDistance := math.Abs(aYCenter-bYCenter) - (a.Height()/2 + b.Height()/2)
	if yDistance < 0 {
		yDistance = 0.0
	}

	return math.Sqrt(math.Pow(xDistance, 2.0) + math.Pow(yDistance, 2.0))
}

// Get approximate distance to coords from object
func (a GameObject) GetDistanceToXY(x float64, y float64) float64 {
	aX := a.X()
	aY := a.Y()

	bX := x
	bY := y

	xDistance := math.Abs(aX - bX)
	yDistance := math.Abs(aY - bY)

	return math.Sqrt(math.Pow(xDistance, 2.0) + math.Pow(yDistance, 2.0))
}

// Determines if 2 objects are close enough to each other
func (a GameObject) IsCloseTo(b IGameObject) bool {
	if a.Floor() != b.Floor() {
		return false
	}
	return a.GetDistance(b) < MaxDistance
}

// Rotates Game object.
func (obj *GameObject) Rotate(rotation float64) {
	if obj.Rotation() != rotation {
		obj.SetRotation(rotation)
		width := obj.Width()
		obj.SetWidth(obj.Height())
		obj.SetHeight(width)
	}
}

// Set rotation depending on the direction
func (obj *GameObject) SetRotationByDirection(direction string) {
	obj.SetRotation(obj.GetRotationByDirection(direction))
}

// Get rotation depending on the direction
func (obj *GameObject) GetRotationByDirection(direction string) float64 {
	validDirection := false
	for _, dir := range constants.GetPossibleDirections() {
		if dir == direction {
			validDirection = true
			break
		}
	}

	if !validDirection {
		//TODO: log error
		return 0.0
	}

	switch direction {
	case "move_north":
		return math.Pi / 2
	case "move_south":
		return math.Pi * 3 / 2
	case "move_east":
		return 0.0
	case "move_west":
		return math.Pi
	case "move_north_east":
		return math.Pi / 4
	case "move_north_west":
		return math.Pi * 3 / 4
	case "move_south_east":
		return math.Pi * 7 / 4
	case "move_south_west":
		return math.Pi * 5 / 4
	}

	return 0.0
}

// Get direction from object to x,y coords
func (obj *GameObject) GetDirectionToXY(x float64, y float64) string {
	possibleDirections := constants.GetPossibleDirections()
	// Calclate angle between mob and target
	// Choose the closest direction by angle by calculatin index in PossibleDirections slice
	dx := x - obj.X()
	dy := y - obj.Y()
	angle := math.Atan2(dy, dx) // range (-PI, PI)
	if angle < 0.0 {
		angle = angle + math.Pi*2
	}
	quotient := math.Floor(angle / (math.Pi / 4)) // math.Pi / 4 - is the angle between movement directions
	remainder := angle - (math.Pi/4)*quotient
	if remainder > math.Pi/8 {
		quotient = quotient + 1.0
	}
	directionIndex := int(quotient)
	if directionIndex == len(possibleDirections) {
		directionIndex = 0
	}
	return possibleDirections[directionIndex]
}

func (obj *GameObject) TurnToXY(x float64, y float64) bool {
	direction := obj.GetDirectionToXY(x, y)
	if obj.Rotation() != obj.GetRotationByDirection(direction) {
		obj.SetRotationByDirection(direction)
		return true
	}
	return false
}
