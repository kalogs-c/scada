package state

import (
	"context"
	"testing"
)

func TestStateMachine_AddAndChangeState_Success(t *testing.T) {
	ctx := context.Background()
	sm := NewStateMachine()

	mock := &mockState{}
	sm.AddState("menu", mock)

	err := sm.Change(ctx, "menu")
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if !mock.entered {
		t.Error("Expected Enter to be called on the new state")
	}

	sm.Update(0.016)
	if !mock.updated {
		t.Error("Expected Update to be called on the current state")
	}

	sm.Render()
	if !mock.rendered {
		t.Error("Expected Render to be called on the current state")
	}
}

func TestStateMachine_ChangeState_InvalidKey(t *testing.T) {
	ctx := context.Background()
	sm := NewStateMachine()

	err := sm.Change(ctx, "invalid")
	if err == nil {
		t.Fatal("Expected error when changing to an unknown state key, got nil")
	}
}

func TestStateMachine_ChangeState_EnterError(t *testing.T) {
	ctx := context.Background()
	sm := NewStateMachine()

	badState := &mockStateWithError{}
	sm.AddState("bad", badState)

	err := sm.Change(ctx, "bad")
	if err == nil {
		t.Fatal("Expected error from state's Enter, got nil")
	}
}
