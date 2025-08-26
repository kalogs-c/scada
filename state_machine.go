package scada

import (
	"context"
	"fmt"
)

// StateMachine manages different states and handles transitions between them.
type StateMachine struct {
	current State
	states  map[string]State
}

// NewStateMachine creates a new StateMachine with optional initial states.
func NewStateMachine(states ...map[string]State) StateMachine {
	empty := emptyState{}
	defaultStates := make(map[string]State)
	if len(states) > 0 {
		defaultStates = states[0]
	}
	return StateMachine{current: empty, states: defaultStates}
}

// AddState registers a new state with the given key.
func (sm *StateMachine) AddState(stateKey string, state State) {
	sm.states[stateKey] = state
}

// Change transitions to a new state by key, calling Exit on the current and Enter on the new state.
func (sm *StateMachine) Change(ctx context.Context, stateKey string) error {
	newState, ok := sm.states[stateKey]
	if !ok {
		return fmt.Errorf("error changing state: state key %s provided does not exists", stateKey)
	}

	sm.current.Exit()
	sm.current = newState
	return sm.current.Enter(ctx)
}

// Update calls the Update method of the current state.
func (sm *StateMachine) Update(dt float32) {
	sm.current.Update(dt)
}

// Render calls the Render method of the current state.
func (sm *StateMachine) Render() {
	sm.current.Render()
}
