package scada

import "context"

// State defines the interface for application states managed by the StateMachine.
// Each state must implement Enter, Update, Render, and Exit methods.
type State interface {
	// Enter is called when the state becomes active.
	Enter(ctx context.Context) error
	// Update is called every frame with the time delta.
	Update(dt float32)
	// Render is called every frame to draw the state.
	Render()
	// Exit is called before the state is deactivated.
	Exit()
}
