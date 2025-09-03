package scada

import (
	"context"
)

// StateMachine manages different states and handles transitions between them.
type StateMachine[T comparable] struct {
	current State
	states  map[T]StateFactory
}

// NewStateMachine creates a new StateMachine with optional initial states.
func NewStateMachine[T comparable]() *StateMachine[T] {
	return &StateMachine[T]{
		current: emptyState{},
		states:  make(map[T]StateFactory),
	}
}

// AddState registers a new state with the given key.
func (sm *StateMachine[T]) AddState(stateKey T, factory StateFactory) {
	sm.states[stateKey] = factory
}

// Change transitions to a new state by key, calling Exit on the current and Enter on the new state.
func (sm *StateMachine[T]) Change(ctx context.Context, stateKey T, data any) {
	factory, ok := sm.states[stateKey]
	if !ok {
		return
	}

	sm.current.Exit()
	sm.current = factory(ctx, data)
}

// Update calls the Update method of the current state.
func (sm *StateMachine[T]) Update(dt float32) {
	sm.current.Update(dt)
}

// Render calls the Render method of the current state.
func (sm *StateMachine[T]) Render() {
	sm.current.Render()
}
