package state

import "context"

type State interface {
	Enter(ctx context.Context) error
	Update(dt float32)
	Render()
}
