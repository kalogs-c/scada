package scada

import "context"

type emptyState struct{}

func newEmptyState() State {
	return emptyState{}
}

func (e emptyState) Enter(ctx context.Context) error {
	return nil
}

func (e emptyState) Update(dt float32) {}
func (e emptyState) Render()           {}
