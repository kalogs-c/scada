package scada

import (
	"context"
	"fmt"
)

type StateMachine struct {
	current State
	states  map[string]State
}

func NewStateMachine(states ...map[string]State) StateMachine {
	empty := emptyState{}
	defaultStates := make(map[string]State)
	if len(states) > 0 {
		defaultStates = states[0]
	}
	return StateMachine{current: empty, states: defaultStates}
}

func (sm *StateMachine) AddState(stateKey string, state State) {
	sm.states[stateKey] = state
}

func (sm *StateMachine) Change(ctx context.Context, stateKey string) error {
	newState, ok := sm.states[stateKey]
	if !ok {
		return fmt.Errorf("error changing state: state key %s provided does not exists", stateKey)
	}

	sm.current.Exit()
	sm.current = newState
	return sm.current.Enter(ctx)
}

func (sm *StateMachine) Update(dt float32) {
	sm.current.Update(dt)
}

func (sm *StateMachine) Render() {
	sm.current.Render()
}
