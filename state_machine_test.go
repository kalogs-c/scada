package scada

import (
	"context"
	"errors"
	"testing"
)

type fakeState struct {
	enterCalled  bool
	exitCalled   bool
	updateCalled bool
	renderCalled bool
	enterErr     error
	ctxReceived  context.Context
}

func (f *fakeState) Enter(ctx context.Context) error {
	f.enterCalled = true
	f.ctxReceived = ctx
	return f.enterErr
}
func (f *fakeState) Exit()             { f.exitCalled = true }
func (f *fakeState) Update(dt float32) { f.updateCalled = true }
func (f *fakeState) Render()           { f.renderCalled = true }

func TestAddStateAndChange(t *testing.T) {
	sm := NewStateMachine()
	state1 := &fakeState{}
	state2 := &fakeState{}
	sm.AddState("s1", state1)
	sm.AddState("s2", state2)

	ctx := context.Background()
	err := sm.Change(ctx, "s1")
	if err != nil {
		t.Errorf("expected no error on first Change, got %v", err)
	}
	if !state1.enterCalled {
		t.Error("Enter should be called on state1")
	}
	// Change to state2, should call Exit on state1 and Enter on state2
	state1.enterCalled = false // reset to check future calls
	err = sm.Change(ctx, "s2")
	if err != nil {
		t.Errorf("expected no error on Change to s2, got %v", err)
	}
	if !state1.exitCalled {
		t.Error("Exit should be called on previous state (state1)")
	}
	if !state2.enterCalled {
		t.Error("Enter should be called on new state (state2)")
	}
}

func TestChangeToNonExistentState(t *testing.T) {
	sm := NewStateMachine()
	err := sm.Change(context.Background(), "missing")
	if err == nil {
		t.Error("expected error when changing to non-existent state")
	}
}

func TestUpdateAndRenderCallsCurrentState(t *testing.T) {
	sm := NewStateMachine()
	state := &fakeState{}
	sm.AddState("main", state)
	_ = sm.Change(context.Background(), "main")

	sm.Update(0.5)
	sm.Render()
	if !state.updateCalled {
		t.Error("Update should call Update on the current state")
	}
	if !state.renderCalled {
		t.Error("Render should call Render on the current state")
	}
}

func TestChangePropagatesEnterError(t *testing.T) {
	sm := NewStateMachine()
	state := &fakeState{enterErr: errors.New("fail")}
	sm.AddState("fail", state)
	err := sm.Change(context.Background(), "fail")
	if err == nil || err.Error() != "fail" {
		t.Error("Change should propagate Enter error")
	}
}
