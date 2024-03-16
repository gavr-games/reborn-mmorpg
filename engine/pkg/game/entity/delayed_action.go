package entity

import (
	"encoding/json"
	"go.uber.org/atomic"
	"sync"
)

const (
	DelayedActionReady = 1
	DelayedActionStarted = 2
)

// Allows engine to perform functions in some time
type DelayedAction struct {
	funcName *atomic.String
	params   map[string]interface{}
	timeLeft *atomic.Float64 // Milliseconds
	status   *atomic.Int64
	mu       sync.RWMutex
}

func NewDelayedAction(funcName string, params map[string]interface{}, timeLeft float64, status int) *DelayedAction {
	return &DelayedAction{
		funcName: atomic.NewString(funcName),
		params: params,
		timeLeft: atomic.NewFloat64(timeLeft),
		status: atomic.NewInt64(int64(status)),
	}
}

func (da *DelayedAction) FuncName() string {
	return da.funcName.Load()
}

func (da *DelayedAction) SetFuncName(funcName string) {
	da.funcName.Store(funcName)
}

func (da *DelayedAction) Params() map[string]interface{} {
	da.mu.RLock()
	defer da.mu.RUnlock()
	return da.params
}

func (da *DelayedAction) SetParams(params map[string]interface{}) {
	da.mu.Lock()
	defer da.mu.Unlock()
	da.params = params
}

func (da *DelayedAction) TimeLeft() float64 {
	return da.timeLeft.Load()
}

func (da *DelayedAction) SetTimeLeft(timeLeft float64) {
	da.timeLeft.Store(timeLeft)
}

func (da *DelayedAction) Status() int {
	return int(da.status.Load())
}

func (da *DelayedAction) SetStatus(status int) {
	da.status.Store(int64(status))
}

func (da *DelayedAction) UnmarshalJSON(b []byte) error {
	var tmp struct {
		FuncName string
		TimeLeft float64
		Params   map[string]interface{}
		Status   int
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}
	da.funcName = atomic.NewString(tmp.FuncName)
	da.timeLeft = atomic.NewFloat64(tmp.TimeLeft)
	da.params   = tmp.Params
	da.status   = atomic.NewInt64(int64(tmp.Status))
	return nil
}

func (da *DelayedAction) MarshalJSON() ([]byte, error) {
	da.mu.RLock()
	defer da.mu.RUnlock()
	return json.Marshal(struct {
		FuncName string
		TimeLeft float64
		Params   map[string]interface{}
		Status   int
	}{
		FuncName: da.FuncName(),
		TimeLeft: da.TimeLeft(),
		Params:   da.Params(),
		Status:   da.Status(),
	})
}
