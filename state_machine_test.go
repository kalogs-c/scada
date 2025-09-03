package scada

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testState struct {
	entered  bool
	exited   bool
	updated  bool
	rendered bool
}

func (s *testState) Update(dt float32) {
	s.updated = true
}

func (s *testState) Render() {
	s.rendered = true
}

func (s *testState) Exit() {
	s.exited = true
}

// factory simples que marca entered
func newTestState() StateFactory {
	return func(ctx context.Context, data any) State {
		s := &testState{}
		s.entered = true
		return s
	}
}

func TestStateMachine_AddChangeState(t *testing.T) {
	sm := NewStateMachine[string]()
	sm.AddState("state1", newTestState())
	sm.AddState("state2", newTestState())

	sm.Change(context.Background(), "state1", nil)
	curr1 := sm.current
	require.NotNil(t, curr1, "current state should not be nil")

	sm.Update(1)
	sm.Render()

	ts1, ok := curr1.(*testState)
	require.True(t, ok, "current state should be a testState")

	assert.True(t, ts1.updated, "Update should be called on state")
	assert.True(t, ts1.rendered, "Render should be called on state")
	assert.True(t, ts1.entered, "state should be marked entered")

	sm.Change(context.Background(), "state2", nil)
	assert.True(t, ts1.exited, "Exit should be called on previous state")

	curr2 := sm.current
	assert.NotEqual(t, curr1, curr2, "current state should change after Change")
}
