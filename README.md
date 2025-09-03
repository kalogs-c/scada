# SCADA – Game Engine Helpers for Go + Raylib

SCADA is a lightweight Go library that provides utilities for building 2D games using [Raylib](https://github.com/gen2brain/raylib-go). It includes a **state machine**, **event bus**, and helper components to organize game logic cleanly and safely.

---

## Features

* **StateMachine** – manage game states with transitions, factories, and context-aware lifecycle.
* **EventBus** – type-safe, concurrent event handling with subscribers, default actions, and automatic context cancellation.
* **Engine** – simplified game loop integrating your states and rendering via Raylib.

---

## Installation

```bash
go get github.com/kalogs-c/scada
```

---

## Usage

### Engine + StateMachine

```go
package main

import (
    "context"
    "github.com/yourusername/scada"
    rl "github.com/gen2brain/raylib-go/raylib"
)

type GameStateKey int

const (
    PlayStateKey GameStateKey = iota
    GameOverStateKey
)

func main() {
    sm := scada.NewStateMachine[GameStateKey]()

    // Add states using factories
    sm.AddState(PlayStateKey, scada.StateFactory(PlayStateFactory))
    sm.AddState(GameOverStateKey, scada.StateFactory(GameOverStateFactory))

    engine := scada.NewEngine(sm)
    engine.Start(800, 600, "My Game", 60)
}
```

* Each state is created on demand using a **StateFactory**.
* The StateMachine uses a **generic StateKey**, so you are not limited to strings.

---

### EventBus

```go
bus := scada.NewEventBus[int]()

ctx, cancel := context.WithCancel(context.Background())
defer cancel()

ch := bus.Sub(ctx, 1, 1, func(e scada.Event[int]) {
    // default action if channel is full
})

bus.Pub(scada.NewEvent(1, 100))
event := <-ch
fmt.Println(event.Payload) // 100
```

* `EventBus[T]` is generic and thread-safe.
* Subscribers automatically unsubscribe when their context is canceled.
* Default actions are executed if the subscriber channel is full.

---

## Design Notes

* **StateMachine** creates states on demand via factories, keeping states decoupled.
* **EventBus** uses generics for type safety and channels with context-aware lifecycle.
* **Engine** integrates the game loop with update/render calls for the current state.
* Helpers are modular and easy to reuse in any 2D Raylib game.

---

## License

MIT License © 2025 Carlos Camilo
