package state

import (
	"context"
	"errors"
)

type mockState struct {
	entered  bool
	updated  bool
	rendered bool
}

func (m *mockState) Enter(ctx context.Context) error {
	m.entered = true
	return nil
}

func (m *mockState) Update(dt float32) {
	m.updated = true
}

func (m *mockState) Render() {
	m.rendered = true
}

type mockStateWithError struct{}

func (m *mockStateWithError) Enter(ctx context.Context) error {
	return errors.New("failed to enter state")
}

func (m *mockStateWithError) Update(dt float32) {}
func (m *mockStateWithError) Render()           {}
