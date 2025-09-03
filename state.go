package scada

import "context"

// What is called to create and enter the state
type StateFactory func(ctx context.Context, data any) State

// State defines the interface for application states managed by the StateMachine.
// Each state must implement Enter, Update, Render, and Exit methods.
type State interface {
	// Update is called every frame with the time delta.
	Update(dt float32)
	// Render is called every frame to draw the state.
	Render()
	// Exit is called before the state is deactivated.
	Exit()
}
