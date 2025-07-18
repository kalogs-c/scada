package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/kalogs-c/scada/state"
)

type Engine struct {
	sm *state.StateMachine
}

func NewEngine(sm *state.StateMachine) Engine {
	return Engine{sm}
}

func (e *Engine) Start(width, height int32, title string, targetFps int32) {
	rl.InitWindow(width, height, title)
	defer rl.CloseWindow()

	rl.SetTargetFPS(targetFps)

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()
		e.sm.Update(dt)

		rl.BeginDrawing()
		e.sm.Render()
		rl.EndDrawing()
	}
}
