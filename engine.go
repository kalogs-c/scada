// Package scada provides a simple state-driven engine using Raylib for rendering.
package scada

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Engine is the main application engine that manages the game loop and state machine.
type Engine struct {
	sm *StateMachine
}

// NewEngine creates a new Engine instance with the given state machine.
func NewEngine(sm *StateMachine) Engine {
	return Engine{sm}
}

// Start initializes the window and runs the main loop using the provided parameters.
// It handles updating and rendering the current state via the state machine.
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
