package entity

import (
	"encoding/json"
	"math"
	"reflect"
	"sync"

	"go.uber.org/atomic"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/constants"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

const (
	MaxDistance = 0.2
)

type IGameObject interface {
	InitGameObject()
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
	GameAreaId() string
	SetGameAreaId(gameAreaId string)
	Rotation() float64
	SetRotation(rotation float64)
	CurrentAction() *DelayedAction
	SetCurrentAction(currentAction *DelayedAction)
	MoveToCoords() *MoveToCoords
	SetMoveToCoords(moveToCoords *MoveToCoords)
	SetMoveToCoordsByObject(moveToObj IGameObject, callback func())
	SetMoveToCoordsByXY(x float64, y float64, callback func())
	Properties() map[string]interface{}
	SetProperties(properties map[string]interface{})
	GetProperty(key string) interface{}
	SetProperty(key string, value interface{})
	Effects() map[string]interface{}
	SetEffects(effects map[string]interface{})
	GetEffect(key string) interface{}
	SetEffect(key string, value interface{})
	RemoveEffect(key string)
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
	x      *atomic.Float64
	y      *atomic.Float64
	width  *atomic.Float64
	height *atomic.Float64

	// game params
	id            *atomic.String
	objType       *atomic.String
	gameAreaId    *atomic.String          // "" - does not belong to any area
	currentAction *atomic.Pointer[DelayedAction]
	rotation      *atomic.Float64         // from 0 to math.Pi * 2
	properties    map[string]interface{}  //TODO: Refactor to thread safe access
	effects       map[string]interface{}  //TODO: Refactor to thread safe access
	moveToCoords  *atomic.Pointer[MoveToCoords]  //used for engine to automatically move object to this coord. TODO: Refactor to thread safe access
	propsMutex    sync.RWMutex
	effectsMutex  sync.RWMutex
}

func (obj *GameObject) InitGameObject() {
	obj.x = atomic.NewFloat64(0.0)
	obj.y = atomic.NewFloat64(0.0)
	obj.width = atomic.NewFloat64(0.0)
	obj.height = atomic.NewFloat64(0.0)
	obj.id = atomic.NewString("")
	obj.objType = atomic.NewString("")
	obj.gameAreaId = atomic.NewString("")
	obj.currentAction = atomic.NewPointer[DelayedAction](nil)
	obj.moveToCoords = atomic.NewPointer[MoveToCoords](nil)
	obj.rotation = atomic.NewFloat64(0.0)
	obj.SetProperties(make(map[string]interface{}))
	obj.SetEffects(make(map[string]interface{}))
}

func (obj *GameObject) X() float64 {
	return obj.x.Load()
}

func (obj *GameObject) SetX(x float64) {
	obj.x.Store(x)
	obj.propsMutex.Lock()
	defer obj.propsMutex.Unlock()
	obj.properties["x"] = x
}

func (obj *GameObject) Y() float64 {
	return obj.y.Load()
}

func (obj *GameObject) SetY(y float64) {
	obj.y.Store(y)
	obj.propsMutex.Lock()
	defer obj.propsMutex.Unlock()
	obj.properties["y"] = y
}

func (obj *GameObject) Width() float64 {
	return obj.width.Load()
}

func (obj *GameObject) SetWidth(width float64) {
	obj.width.Store(width)
	obj.propsMutex.Lock()
	defer obj.propsMutex.Unlock()
	obj.properties["width"] = width
}

func (obj *GameObject) Height() float64 {
	return obj.height.Load()
}

func (obj *GameObject) SetHeight(height float64) {
	obj.height.Store(height)
	obj.propsMutex.Lock()
	defer obj.propsMutex.Unlock()
	obj.properties["height"] = height
}

func (obj *GameObject) Id() string {
	return obj.id.Load()
}

func (obj *GameObject) SetId(id string) {
	obj.id.Store(id)
	obj.propsMutex.Lock()
	defer obj.propsMutex.Unlock()
	obj.properties["id"] = id
}

func (obj *GameObject) Kind() string {
	obj.propsMutex.RLock()
	defer obj.propsMutex.RUnlock()
	return obj.properties["kind"].(string)
}

func (obj *GameObject) Type() string {
	return obj.objType.Load()
}

func (obj *GameObject) SetType(t string) {
	obj.objType.Store(t)
	obj.propsMutex.Lock()
	defer obj.propsMutex.Unlock()
	obj.properties["type"] = t
}

func (obj *GameObject) GameAreaId() string {
	return obj.gameAreaId.Load()
}

func (obj *GameObject) SetGameAreaId(gameAreaId string) {
	obj.gameAreaId.Store(gameAreaId)
}

func (obj *GameObject) Rotation() float64 {
	return obj.rotation.Load()
}

func (obj *GameObject) SetRotation(rotation float64) {
	obj.rotation.Store(rotation)
}

func (obj *GameObject) CurrentAction() *DelayedAction {
	return obj.currentAction.Load()
}

func (obj *GameObject) SetCurrentAction(currentAction *DelayedAction) {
	obj.currentAction.Store(currentAction)
}

func (obj *GameObject) Properties() map[string]interface{} {
	obj.propsMutex.RLock()
	defer obj.propsMutex.RUnlock()
	return utils.CopyMap(obj.properties) // returns a copy of map, so external code could work without Mutex Lock (no parallel access to original properties map)
}

func (obj *GameObject) SetProperties(properties map[string]interface{}) {
	obj.propsMutex.Lock()
	defer obj.propsMutex.Unlock()
	obj.properties = properties
}

func (obj *GameObject) GetProperty(key string) interface{} {
	obj.propsMutex.RLock()
	defer obj.propsMutex.RUnlock()
	value := obj.properties[key]
	if value == nil {
		return nil
	}
	// We want to create copies of maps and slices, so they stay "immutable" and allow safe parallel access
	switch reflect.TypeOf(value).Kind() {
	case reflect.Map:
		return utils.CopyMap(value.(map[string]interface{}))
	case reflect.Slice:
		var slice []interface{} = make([]interface{}, len(value.([]interface{})))
		for k, v := range value.([]interface{}) {
			slice[k] = v
		}
		return slice
	default:
		return value
	}
}

func (obj *GameObject) SetProperty(key string, value interface{}) {
	obj.propsMutex.Lock()
	defer obj.propsMutex.Unlock()
	obj.properties[key] = value
}

func (obj *GameObject) Effects() map[string]interface{} {
	obj.effectsMutex.RLock()
	defer obj.effectsMutex.RUnlock()
	return utils.CopyMap(obj.effects) // returns a copy of map, so external code could work without Mutex Lock (no parallel access to original effects map)
}

func (obj *GameObject) SetEffects(effects map[string]interface{}) {
	obj.effectsMutex.Lock()
	defer obj.effectsMutex.Unlock()
	obj.effects = effects
}

func (obj *GameObject) GetEffect(key string) interface{} {
	obj.effectsMutex.RLock()
	defer obj.effectsMutex.RUnlock()
	return obj.effects[key]
}

func (obj *GameObject) SetEffect(key string, value interface{}) {
	obj.effectsMutex.Lock()
	defer obj.effectsMutex.Unlock()
	obj.effects[key] = value
}

func (obj *GameObject) RemoveEffect(key string) {
	obj.effectsMutex.Lock()
	defer obj.effectsMutex.Unlock()
	obj.effects[key] = nil
	delete(obj.effects, key)
}

func (obj *GameObject) MoveToCoords() *MoveToCoords {
	return obj.moveToCoords.Load()
}

func (obj *GameObject) SetMoveToCoords(moveToCoords *MoveToCoords) {
	obj.moveToCoords.Store(moveToCoords)
}

func (obj *GameObject) SetMoveToCoordsByObject(moveToObj IGameObject, callback func()) {
	obj.SetMoveToCoords(&MoveToCoords{
		Mode: MoveCloseToBounds,
		Bounds: utils.Bounds{
			X:      moveToObj.X(),
			Y:      moveToObj.Y(),
			Width:  moveToObj.Width(),
			Height: moveToObj.Height(),
		},
		DirectionChangeTime:      constants.MoveToDefaultDirectionChangeTime,
		TimeUntilDirectionChange: 0,
		Callback: callback,
	})
}

func (obj *GameObject) SetMoveToCoordsByXY(x float64, y float64, callback func()) {
	obj.SetMoveToCoords(&MoveToCoords{
		Mode: MoveToExactCoords,
		Bounds: utils.Bounds{
			X:      x,
			Y:      y,
			Width:  0.0,
			Height: 0.0,
		},
		DirectionChangeTime:      constants.MoveToDefaultDirectionChangeTime,
		TimeUntilDirectionChange: 0,
		Callback: callback,
	})
}

func (obj *GameObject) UnmarshalJSON(b []byte) error {
	//TODO: Unmarshal moveToCoords
	var tmp struct {
		X             float64
		Y             float64
		Width         float64
		Height        float64
		Id            string
		Type          string
		GameAreaId    string
		CurrentAction *DelayedAction
		Rotation      float64
		Properties    map[string]interface{}
		Effects       map[string]interface{}
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}
	obj.InitGameObject()
	obj.SetX(tmp.X)
	obj.SetY(tmp.Y)
	obj.SetWidth(tmp.Width)
	obj.SetHeight(tmp.Height)
	obj.SetId(tmp.Id)
	obj.SetType(tmp.Type)
	obj.SetGameAreaId(tmp.GameAreaId)
	obj.SetCurrentAction(tmp.CurrentAction)
	obj.SetRotation(tmp.Rotation)
	obj.SetProperties(tmp.Properties)
	obj.SetEffects(tmp.Effects)
	return nil
}

func (obj *GameObject) MarshalJSON() ([]byte, error) {
	//TODO: Marshal moveToCoords
	return json.Marshal(struct {
		X             float64
		Y             float64
		Width         float64
		Height        float64
		Id            string
		Type          string
		GameAreaId    string
		CurrentAction *DelayedAction
		Rotation      float64
		Properties    map[string]interface{}
		Effects       map[string]interface{}
	}{
		X:             obj.X(),
		Y:             obj.Y(),
		Width:         obj.Width(),
		Height:        obj.Height(),
		Id:            obj.Id(),
		Type:          obj.Type(),
		GameAreaId:    obj.GameAreaId(),
		CurrentAction: obj.CurrentAction(),
		Rotation:      obj.Rotation(),
		Properties:    obj.Properties(),
		Effects:       obj.Effects(),
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
	//TODO: Clone currentAction and moveToCoords
	clone := &GameObject{
		x:             atomic.NewFloat64(obj.X()),
		y:             atomic.NewFloat64(obj.Y()),
		width:         atomic.NewFloat64(obj.Width()),
		height:        atomic.NewFloat64(obj.Height()),
		id:            atomic.NewString(obj.Id()),
		objType:       atomic.NewString(obj.Type()),
		gameAreaId:    atomic.NewString(obj.GameAreaId()),
		rotation:      atomic.NewFloat64(obj.Rotation()),
		currentAction: atomic.NewPointer[DelayedAction](nil),
		moveToCoords:  atomic.NewPointer[MoveToCoords](nil),
		properties:    make(map[string]interface{}),
		effects:       make(map[string]interface{}),
	}
	clone.SetProperties(obj.Properties())
	clone.SetEffects(obj.Effects())
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
	if a.GameAreaId() != b.GameAreaId() {
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
