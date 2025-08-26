package scada

import "context"

// emptyState is a no-op implementation of the State interface,
// used as a default or placeholder state.
type emptyState struct{}

// Enter does nothing and returns nil for emptyState.
func (e emptyState) Enter(ctx context.Context) error {
	return nil
}

// Update does nothing for emptyState.
func (e emptyState) Update(dt float32) {}

// Render does nothing for emptyState.
func (e emptyState) Render() {}

// Exit does nothing for emptyState.
func (e emptyState) Exit() {}
