package entity

// Allows engine to perform functions in some time
type DelayedAction struct {
	FuncName string
	Params map[string]interface{}
	TimeLeft float64 // Milliseconds
}
