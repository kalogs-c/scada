package state

import (
	"context"
	"fmt"
)

type StateMachine struct {
	current State
	states  map[string]State
}

func NewStateMachine() StateMachine {
	empty := newEmptyState()
	states := make(map[string]State)
	return StateMachine{current: empty, states: states}
}

func (sm *StateMachine) AddState(stateKey string, state State) {
	sm.states[stateKey] = state
}

func (sm *StateMachine) Change(ctx context.Context, stateKey string) error {
	newState, ok := sm.states[stateKey]
	if !ok {
		return fmt.Errorf("error changing state: state key provided does not exists. state key: %s", stateKey)
	}

	sm.current = newState
	return sm.current.Enter(ctx)
}

func (sm *StateMachine) Update(dt float32) {
	sm.current.Update(dt)
}

func (sm *StateMachine) Render() {
	sm.current.Render()
}
